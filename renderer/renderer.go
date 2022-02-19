package renderer

import (
	"fmt"
	"log"

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
