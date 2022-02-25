package main

import (
	"fmt"
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
	"github.com/engoengine/glm"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width = 960
	height = 540
	cpx = width/2.0
	cpy = height/2.0
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

	window, err := glfw.CreateWindow(width, height, "Draw a Square: Projection Matrix", nil, nil)
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
	
	sqDim := 350.0
	sqOffset := float32(sqDim/2.0)
	positions := []float32{ // use a slice
		// texture coordinates added -- may need to flip here if upside down
		cpx - sqOffset, cpy + sqOffset, 0.0, 0.0,	// vert TL - index 0
		cpx - sqOffset, cpy - sqOffset, 0.0, 1.0,	// vert BL - index 1
		cpx + sqOffset, cpy - sqOffset,	1.0, 1.0,	// vert BR - index 2
		cpx + sqOffset, cpy + sqOffset,	1.0, 0.0,	// vert TR - index 3
	}

	indices := []uint32{
		0, 1, 2,
		2, 3, 0,
	}

	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

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

	proj := glm.Ortho(1.0, float32(width), 1.0, float32(height), -1.0, 1.0)
	view := glm.Translate3D(-100.0, 0.0, 0.0)
	model := glm.Translate3D(100.0, 100.0, 0.0)
	mv := proj.Mul4(&view)
	mvp := mv.Mul4(&model)
	
	shader := shader.New("basic.shader")
	defer shader.Close()
	
	tx := texture.New("mimp.png")
	
	var txSlot int32 = 0
	tx.Bind(txSlot)
	uniform_texture := shaderUniform.New(shader, "u_Texture")
	uniform_texture.SetUniform1i(txSlot)
	
	uniform_mvp := shaderUniform.New(shader, "u_MVP")
	uniform_mvp.SetUniformMatrix4fv(mvp)

	for !window.ShouldClose() {

		renderer.Clear()
		renderer.Draw(va, ib, shader)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
