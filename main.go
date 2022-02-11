package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
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
		-0.5, 0.5,
		0.0, -0.5,
		0.5, -0.5,
	}

	// Create our buffer
	var buffer uint32
	gl.GenBuffers(1, &buffer);
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions) * floatSize, gl.Ptr(positions), gl.STATIC_DRAW)

	// this must be called _after_ gl.BindBuffer()
	var vertexIndex uint32 = 0
	var floatsPerAttrib int32 = 2
	gl.EnableVertexAttribArray(vertexIndex)
	gl.VertexAttribPointer(vertexIndex, floatsPerAttrib, gl.FLOAT, false, floatsPerAttrib * int32(floatSize), gl.Ptr(0))

	for !window.ShouldClose() {

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}