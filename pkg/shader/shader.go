package shader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	RendererID uint32
	FilePath   string
	// caching for uniforms
}

type shaderData struct {
	Source string
	Type   uint32
	Name   string
}

type shaderParserData struct {
	VertexShader   shaderData
	FragmentShader shaderData
}

func New(shaderFile string) *Shader {
	s := Shader{}
	s.RendererID = 0
	err := s.locateShader(shaderFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("DEBUG: Shader file located: %s", s.FilePath))
	shaderSource := s.parseShader()
	// log.Println(fmt.Sprintf("DEBUG: Vertex shader:\n%v", shaderSource.VertexShader))
	// log.Println(fmt.Sprintf("DEBUG: Fragment shader:\n%v", shaderSource.FragmentShader))
	s.createShaders(shaderSource)
	log.Println(fmt.Sprintf("DEBUG: Program ID = %d", s.RendererID))
	s.Bind()

	log.Println(fmt.Sprintf("Shaders loaded from %s", s.FilePath))
	return &s
}

func (s *Shader) Close() {
	gl.DeleteProgram(s.RendererID)
}

func (s *Shader) Bind() {
	gl.UseProgram(s.RendererID)
}

func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

func (s *Shader) locateShader(shaderFile string) error {
	s.FilePath = "res/shaders/" + shaderFile
	if _, err := os.Stat(s.FilePath); err != nil {
		return fmt.Errorf("Shader definition '%s' not found", s.FilePath)
	}
	return nil
}

func (s *Shader) parseShader() shaderParserData {
	type ShaderType int
	const (
		None     ShaderType = -1
		Vertex   ShaderType = 0
		Fragment ShaderType = 1
	)

	file, err := os.Open(s.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var shader [2]shaderData
	currentType := None

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "// shader: vertex" {
			currentType = Vertex
			shader[currentType].Name = "Vertex"
			shader[currentType].Type = gl.VERTEX_SHADER
			shader[currentType].Source = ""
		} else if line == "// shader: fragment" {
			currentType = Fragment
			shader[currentType].Name = "Fragment"
			shader[currentType].Type = gl.FRAGMENT_SHADER
			shader[currentType].Source = ""
		} else {
			if currentType != None {
				shader[currentType].Source += line + "\n"
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	shader[Vertex].Source += "\x00"
	shader[Fragment].Source += "\x00"

	var shaders shaderParserData
	shaders.VertexShader = shader[Vertex]
	shaders.FragmentShader = shader[Fragment]
	return shaders
}

func (s *Shader) createShaders(shaderSource shaderParserData) {
	log.Println("DEBUG: in createShaders()")
	programId := gl.CreateProgram()

	log.Println("DEBUG: Compile Vertex Shader")
	vsId, err := s.compileShader(shaderSource.VertexShader)
	if err != nil {
		panic(err)
	}
	
	log.Println("DEBUG: Compile Fragment Shader")
	fsId, err := s.compileShader(shaderSource.FragmentShader)
	if err != nil {
		panic(err)
	}

	log.Println("DEBUG: Attach shaders")
	gl.AttachShader(programId, vsId)
	gl.AttachShader(programId, fsId)
	log.Println("DEBUG: Link")
	gl.LinkProgram(programId)
	log.Println("DEBUG: Validate")
	gl.ValidateProgram(programId)

	gl.DeleteShader(vsId)
	gl.DeleteShader(fsId)

	s.RendererID = programId
}

func (s *Shader) compileShader(shader shaderData) (uint32, error) {
	shaderId := gl.CreateShader(shader.Type)

	sourcePtr, free := gl.Strs(shader.Source)
	gl.ShaderSource(shaderId, 1, sourcePtr, nil)
	free()
	gl.CompileShader(shaderId)

	// Error handling
	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)
		message := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, &logLength, gl.Str(message))

		// standard Go would return an error here, but for this tutorial
		// we shall simply print it out instead
		gl.DeleteShader(shaderId)
		return 0, fmt.Errorf("shader '%s' in %s has not compiled: %v", shader.Name, s.FilePath, message)
	}

	return shaderId, nil
}
