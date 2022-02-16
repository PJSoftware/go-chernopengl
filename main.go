package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

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

	window, err := glfw.CreateWindow(640, 480, "Draw a Square", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Important! Call gl.Init only under the presence of an active OpenGL context,
	// i.e., after MakeContextCurrent.
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	log.Println(gl.VERSION)
	
	floatSize := 4	// a float32 is 32 bits, or 4 bytes, in size
	positions := []float32{ // use a slice
		-0.5,  0.5,		// vert TL - index 0
		-0.5, -0.5,		// vert BL - index 1
		 0.5, -0.5,		// vert BR - index 2
		 0.5,  0.5,		// vert TR - index 3
	}

	intSize := 4	// a uint32 is 32 bits
	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	// Create our vertex buffer
	var vertexBuffer uint32
	var numVBuffers int32 = 1
	gl.GenBuffers(numVBuffers, &vertexBuffer);
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions) * floatSize, gl.Ptr(positions), gl.STATIC_DRAW)

	// this must be called _after_ gl.BindBuffer()
	var vertexIndex uint32 = 0
	var floatsPerAttrib int32 = 2
	gl.EnableVertexAttribArray(vertexIndex)
	gl.VertexAttribPointer(vertexIndex, floatsPerAttrib, gl.FLOAT, false, floatsPerAttrib * int32(floatSize), nil)

	// Create our index indexBuffer
	var indexBuffer uint32
	var numIBuffers int32 = 1
	gl.GenBuffers(numIBuffers, &indexBuffer);
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices) * intSize, gl.Ptr(indices), gl.STATIC_DRAW)
	
	shaderSource := parseShader("basic.shader")
	shader := createShaders(shaderSource)
	gl.UseProgram(shader)

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	gl.DeleteProgram(shader)
}