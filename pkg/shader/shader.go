package shader

type Shader struct {
	RendererID uint32
	FilePath   string
	// caching for uniforms
}

func New(shaderFile string) *Shader {
	s := Shader{}
	s.FilePath = shaderFile

	return &s
}

func (s *Shader) Close() {
}

func (s *Shader) Bind() {
	// gl.BindBuffer(gl.ARRAY_BUFFER, vb.RendererID)
}

func (s *Shader) Unbind() {
	// gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// set uniforms
func (s *Shader) SetUniform4f(name string, f1, f2, f3, f4 float32) {
	// gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// private internal functions

func (s *Shader) GetUniformLocation(name string) uint32 {
	return 0
}