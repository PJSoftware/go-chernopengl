package shaderUniform

import (
	"fmt"
	"log"

	"github.com/PJSoftware/go-chernopengl/pkg/shader"
	"github.com/engoengine/glm"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderUniform struct {
	Name     string
	Location int32
}

func New(shader *shader.Shader, name string) *ShaderUniform {
	su := ShaderUniform{}
	su.Name = name

	location := gl.GetUniformLocation(shader.RendererID, gl.Str(su.Name + "\x00"))
	if location == -1 {
		panic(fmt.Sprintf("Could not locate uniform '%s'", name))
	}
	su.Location = location
	log.Println(fmt.Sprintf("Uniform '%s' location = %d", su.Name, su.Location))
	return &su
}

func (su *ShaderUniform) SetUniform1i(i1 int32) {
	gl.Uniform1i(su.Location, i1)
}

func (su *ShaderUniform) SetUniform4f(f1, f2, f3, f4 float32) {
	gl.Uniform4f(su.Location, f1, f2, f3, f4)
}

func (su *ShaderUniform) SetUniformMatrix4fv(matrix glm.Mat4) {
	gl.UniformMatrix4fv(su.Location, 1, false, &matrix[0])
}
