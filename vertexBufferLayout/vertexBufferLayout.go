package vertexBufferLayout

import "github.com/go-gl/gl/v4.1-core/gl"

var sizeOf = map[uint32]int32{ 
	gl.FLOAT: 				4,
	gl.UNSIGNED_INT: 	4,
	gl.UNSIGNED_BYTE: 1,
}

type VertexBufferElement struct {
	Type       uint32
	Count      int32
	Normalised bool
}

func (vbe *VertexBufferElement) GetSizeOfType() int32 {
	return sizeOf[vbe.Type]
}

type VertexBufferLayout struct {
	Elements []VertexBufferElement
	Stride int32
}

func New() *VertexBufferLayout {
	vbl := VertexBufferLayout{}
	vbl.Stride = 0;
	return &vbl
}

// func (vbl *VertexBufferLayout) Close() {
// 	// gl.DeleteBuffers(1, &va.RendererID)
// }

func (vbl *VertexBufferLayout) Push(eType uint32, count int32) {
	normalised := eType == gl.UNSIGNED_BYTE
	element := VertexBufferElement{ eType, count, normalised }
	vbl.Elements = append(vbl.Elements, element)
	vbl.Stride += element.GetSizeOfType() * count
}

func (vbl *VertexBufferLayout) GetElements() *[]VertexBufferElement {
	return &vbl.Elements
}

func (vbl *VertexBufferLayout) GetStride() int32 {
	return vbl.Stride
}