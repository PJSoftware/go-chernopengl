package vertexArray

import (
	"fmt"
	"log"

	"github.com/PJSoftware/go-chernopengl/vertexBuffer"
	"github.com/PJSoftware/go-chernopengl/vertexBufferLayout"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexArray struct {
	RendererID uint32
}

func New() *VertexArray {
	va := VertexArray{}

	gl.GenVertexArrays(1, &va.RendererID);
	log.Println(fmt.Sprintf("Vertex Array Object %d added", va.RendererID))
	return &va
}

func (va *VertexArray) Close() {
	gl.DeleteVertexArrays(1, &va.RendererID)
}

func (va *VertexArray) AddBuffer(vb *vertexBuffer.VertexBuffer, layout *vertexBufferLayout.VertexBufferLayout) {
	va.Bind()
	vb.Bind()

	var offset int32 = 0 
	elements := layout.GetElements()
	for i, element := range *elements {
		idx := uint32(i)
		gl.EnableVertexAttribArray(idx)
		gl.VertexAttribPointer(idx, element.Count, element.Type, element.Normalised, layout.GetStride(), gl.Ptr(offset))
		offset += element.Count * element.GetSizeOfType()
	}	
}

func (va *VertexArray) Bind() {
	gl.BindVertexArray(va.RendererID)
}

func (va *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}
