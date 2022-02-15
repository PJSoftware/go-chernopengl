package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func importShader(shaderFile string) string {
	content, err := os.ReadFile("shaders/" + shaderFile + ".glsl")
	if err != nil {
		log.Fatal(err)
	}
	return string(content) + "\x00" // shader string must be null-terminated to compile
}

func compileShader(shaderType uint32, sourceFile string) (uint32, error) {
	shaderId := gl.CreateShader(shaderType)
	
	source := importShader(sourceFile)
	sourcePtr, free := gl.Strs(source)
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
		return 0, fmt.Errorf("shader '%v' has not compiled: %v", sourceFile, message)
	}

	return shaderId, nil
}

func createShaders(vertexShaderFile string, fragmentShaderFile string) uint32 {
	programId := gl.CreateProgram()

	vsId, err := compileShader(gl.VERTEX_SHADER, vertexShaderFile)
	if err != nil {
		panic(err)
	}

	fsId, err := compileShader(gl.FRAGMENT_SHADER, fragmentShaderFile)
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

	window, err := glfw.CreateWindow(640, 480, "Hello World", nil, nil)
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
		 0.0,  0.5,
		-0.5, -0.5,
		 0.5, -0.5,
	}

	// Create our buffer
	var buffer uint32
	var numBuffers int32 = 1
	gl.GenBuffers(numBuffers, &buffer);
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions) * floatSize, gl.Ptr(positions), gl.STATIC_DRAW)

	// this must be called _after_ gl.BindBuffer()
	var vertexIndex uint32 = 0
	var floatsPerAttrib int32 = 2
	gl.EnableVertexAttribArray(vertexIndex)
	gl.VertexAttribPointer(vertexIndex, floatsPerAttrib, gl.FLOAT, false, floatsPerAttrib * int32(floatSize), nil)

	shader := createShaders("vertexShader", "fragmentShader")
	gl.UseProgram(shader)

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	gl.DeleteProgram(shader)
}