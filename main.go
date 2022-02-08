package main

import (
	"runtime"

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

	for !window.ShouldClose() {

		// gl.Clear(gl.COLOR_BUFFER_BIT)
		// gl.Begin(gl.TRIANGLES)
		// gl.Vertex2f(-0.5, 0.5)
		// gl.Vertex2f(0.0, -0.5)
		// gl.Vertex2f(0.5, -0.5)
		// gl.End()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}