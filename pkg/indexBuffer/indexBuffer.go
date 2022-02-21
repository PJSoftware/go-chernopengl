package indexBuffer

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type IndexBuffer struct {
	RendererID uint32
	Count int32
}

func New(data interface{}, count int) *IndexBuffer {
	ib := IndexBuffer{}

	gl.GenBuffers(1, &ib.RendererID);
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.RendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(count * 32), gl.Ptr(data), gl.STATIC_DRAW)
	ib.Count = int32(count)
	log.Println(fmt.Sprintf("Index Buffer %d added", ib.RendererID))

	return &ib
}

func (ib *IndexBuffer) Close() {
	gl.DeleteBuffers(1, &ib.RendererID)
}

func (ib *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.RendererID)
}

func (ib *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (ib *IndexBuffer) GetCount() int32 {
	return ib.Count
}