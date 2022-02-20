package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/PJSoftware/go-chernopengl/indexBuffer"
	"github.com/PJSoftware/go-chernopengl/renderer"
	"github.com/PJSoftware/go-chernopengl/vertexArray"
	"github.com/PJSoftware/go-chernopengl/vertexBuffer"
	"github.com/PJSoftware/go-chernopengl/vertexBufferLayout"
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

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func parseShader(shaderFile string) ShaderParserData {
	type ShaderType int
	const (
		None 			ShaderType = -1
		Vertex 		ShaderType = 0
		Fragment 	ShaderType = 1
	)
	const shaderPath = "res/shaders"

	file, err := os.Open(shaderPath + "/" + shaderFile)
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

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

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
	
	floatInBytes := 4	// a float32 is 32 bits, or 4 bytes, in size
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

	// // Create our vertex array object
	// var vao uint32
	// gl.GenVertexArrays(1, &vao)
	// gl.BindVertexArray(vao)

	va := vertexArray.New()
	defer va.Close()

	vb := vertexBuffer.New(positions, len(positions) * floatInBytes)
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

	// gl.BindVertexArray(0)
	// gl.UseProgram(0)
	// vb.Unbind()
	// ib.Unbind()

	var r float32 = 0.0
	var increment float32 = 0.02
	for !window.ShouldClose() {
		
		gl.Clear(gl.COLOR_BUFFER_BIT)
		
		gl.UseProgram(shader)
		gl.Uniform4f(location, r, 0.1, 0.3, 1.0)

		// vb.Bind()
		// gl.BindVertexArray(vao)
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