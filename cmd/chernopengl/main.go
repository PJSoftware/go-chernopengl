package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/PJSoftware/go-chernopengl/pkg/indexBuffer"
	"github.com/PJSoftware/go-chernopengl/pkg/lookup"
	"github.com/PJSoftware/go-chernopengl/pkg/renderer"
	"github.com/PJSoftware/go-chernopengl/pkg/shader"
	"github.com/PJSoftware/go-chernopengl/pkg/shaderUniform"
	"github.com/PJSoftware/go-chernopengl/pkg/texture"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexArray"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexBuffer"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexBufferLayout"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

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

func init() {
	// see https://www.socketloop.com/tutorials/golang-how-to-read-jpg-jpeg-gif-and-png-files
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
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
		// texture coordinates added -- may need to flip here if upside down
		-0.5,  0.5,	0.0, 1.0,	// vert TL - index 0
		-0.5, -0.5, 0.0, 0.0,	// vert BL - index 1
		 0.5, -0.5,	1.0, 0.0,	// vert BR - index 2
		 0.5,  0.5,	1.0, 1.0,	// vert TR - index 3
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
	layout.Push(gl.FLOAT, 2)
	va.AddBuffer(vb, layout)

	ib := indexBuffer.New(indices, len(indices))
	defer ib.Close()

	shader := shader.New("basic.shader")
	defer shader.Close()
	
	tx := texture.New("mimp.jpg")
	var txSlot int32 = 0
	tx.Bind(txSlot)
	uniform_texture := shaderUniform.New(shader, "u_Texture")
	uniform_texture.SetUniform1i(txSlot)

	for !window.ShouldClose() {

		renderer.Clear()
		renderer.Draw(va, ib, shader)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
