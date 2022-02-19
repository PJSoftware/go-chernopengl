package vertexBuffer

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexBuffer struct {
	RendererID uint32
}

func New(data interface{}, size int) *VertexBuffer {
	vb := VertexBuffer{}

	gl.GenBuffers(1, &vb.RendererID);
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.RendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	log.Println(fmt.Sprintf("Vertex Buffer %d added", vb.RendererID))
	return &vb
}

func (vb *VertexBuffer) Close() {
	gl.DeleteBuffers(1, &vb.RendererID)
}

func (vb *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.RendererID)
}

func (vb *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
