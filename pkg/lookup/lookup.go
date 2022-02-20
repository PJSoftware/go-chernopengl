package lookup

import "github.com/go-gl/gl/v4.1-core/gl"

var SizeOf = map[uint32]int32{
	gl.FLOAT:         4,
	gl.UNSIGNED_INT:  4,
	gl.UNSIGNED_BYTE: 1,
}
