package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/PJSoftware/go-chernopengl/pkg/indexBuffer"
	"github.com/PJSoftware/go-chernopengl/pkg/lookup"
	"github.com/PJSoftware/go-chernopengl/pkg/renderer"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexArray"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexBuffer"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexBufferLayout"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type ShaderData struct {
	Source string
	File string
	Type uint32
	Name string
}

type ShaderParserData struct {
	VertexShader ShaderData
	FragmentShader ShaderData
}

func parseShader(shaderFile string) ShaderParserData {
	type ShaderType int
	const (
		None 			ShaderType = -1
		Vertex 		ShaderType = 0
		Fragment 	ShaderType = 1
	)

	shaderPath := "res/shaders/" + shaderFile
	fmt.Println(shaderPath)

	file, err := os.Open(shaderPath)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	var shader [2]ShaderData
	currentType := None

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "// shader: vertex" {
			currentType = Vertex
			shader[currentType].File = shaderFile
			shader[currentType].Name = "Vertex"
			shader[currentType].Type = gl.VERTEX_SHADER
			shader[currentType].Source = ""
		} else if line == "// shader: fragment" {
			currentType = Fragment
			shader[currentType].File = shaderFile
			shader[currentType].Name = "Fragment"
			shader[currentType].Type = gl.FRAGMENT_SHADER
			shader[currentType].Source = ""
		} else {
			if currentType != None {
				shader[currentType].Source += line + "\n"
			}
		}
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	shader[Vertex].Source += "\x00"
	shader[Fragment].Source += "\x00"

	var shaders ShaderParserData
	shaders.VertexShader = shader[Vertex]
	shaders.FragmentShader = shader[Fragment]
	return shaders
}

func compileShader(shader ShaderData) (uint32, error) {
	shaderId := gl.CreateShader(shader.Type)
	
	sourcePtr, free := gl.Strs(shader.Source)
	gl.ShaderSource(shaderId, 1, sourcePtr, nil)
	free()
	gl.CompileShader(shaderId)

	// Error handling
	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)
		message := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, &logLength, gl.Str(message))

		// standard Go would return an error here, but for this tutorial
		// we shall simply print it out instead
		gl.DeleteShader(shaderId)
		return 0, fmt.Errorf("shader '%s' in %s has not compiled: %v", shader.Name, shader.File, message)
	}

	return shaderId, nil
}

func createShaders(shaderSource ShaderParserData) uint32 {
	programId := gl.CreateProgram()

	vsId, err := compileShader(shaderSource.VertexShader)
	if err != nil {
		panic(err)
	}

	fsId, err := compileShader(shaderSource.FragmentShader)
	if err != nil {
		panic(err)
	}
	
	gl.AttachShader(programId, vsId)
	gl.AttachShader(programId, fsId)
	gl.LinkProgram(programId)
	gl.ValidateProgram(programId)

	// Once linked, the standalone shaders can be safely deleted
	gl.DeleteShader(vsId)
	gl.DeleteShader(fsId)

	return programId
}

func setWorkingFolder() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)
	
	if err := os.Chdir(exPath); err != nil {
		return fmt.Errorf("Error:\tCould not move into the directory (%s)", exPath)
	}

	return nil
}

func main() {
	var err error
	err = setWorkingFolder()
	if err != nil {
		panic(err)
	}
	runtime.LockOSThread()
  
	err = glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(640, 480, "Draw a Square: Abstracting", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1) // enable vsync

	// Important! Call gl.Init only under the presence of an active OpenGL context,
	// i.e., after MakeContextCurrent.
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	log.Println(fmt.Sprintf("Initialise OpenGL version %d", gl.VERSION))
	
	positions := []float32{ // use a slice
		-0.5,  0.5,		// vert TL - index 0
		-0.5, -0.5,		// vert BL - index 1
		 0.5, -0.5,		// vert BR - index 2
		 0.5,  0.5,		// vert TR - index 3
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	va := vertexArray.New()
	defer va.Close()

	vb := vertexBuffer.New(positions, len(positions) * int(lookup.SizeOf[gl.FLOAT]))
	defer vb.Close()

	layout := vertexBufferLayout.New()
	layout.Push(gl.FLOAT, 2)
	va.AddBuffer(vb, layout)

	ib := indexBuffer.New(indices, len(indices))
	defer ib.Close()

	shaderSource := parseShader("basic.shader")
	shader := createShaders(shaderSource)
	gl.UseProgram(shader)

	location := gl.GetUniformLocation(shader, gl.Str("u_Colour" + "\x00"))
	if (location == -1) {
		panic("Could not locate uniform 'u_Colour'")
	}

	var r float32 = 0.0
	var increment float32 = 0.02
	for !window.ShouldClose() {
		
		gl.Clear(gl.COLOR_BUFFER_BIT)
		
		gl.UseProgram(shader)
		gl.Uniform4f(location, r, 0.1, 0.3, 1.0)

		va.Bind()
		ib.Bind()
	
		renderer.ClearError()
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		renderer.PanicOnError()
		
		window.SwapBuffers()
		glfw.PollEvents()

		if r >= 1.0 {
			increment = -0.02
		} else if r <= 0.0 {
			increment = 0.02
		} 
		r += increment
	}

	gl.DeleteProgram(shader)
}