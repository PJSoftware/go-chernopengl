package renderer

import (
	"fmt"
	"log"

	"github.com/PJSoftware/go-chernopengl/pkg/indexBuffer"
	"github.com/PJSoftware/go-chernopengl/pkg/shader"
	"github.com/PJSoftware/go-chernopengl/pkg/vertexArray"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func ClearError() {
	for gl.GetError() != gl.NO_ERROR {
	}
}

func PanicOnError() {
	errorOccurred := false

	for {
		glError := gl.GetError()
		if glError == gl.NO_ERROR {
			break
		}
		log.Println(fmt.Sprintf("OpenGL Error #%d", glError))
		errorOccurred = true
	}

	if errorOccurred {
		panic("OpenGL Error(s) detected")
	}
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Draw(va *vertexArray.VertexArray, ib *indexBuffer.IndexBuffer, shader *shader.Shader) {
	shader.Bind()
	va.Bind()
	ib.Bind()
	gl.DrawElements(gl.TRIANGLES, ib.GetCount(), gl.UNSIGNED_INT, nil)
}