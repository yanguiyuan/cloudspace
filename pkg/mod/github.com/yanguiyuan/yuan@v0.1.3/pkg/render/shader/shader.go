package shader

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

func Compile(src string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource, free := gl.Strs(src)
	gl.ShaderSource(shader, 1, cSource, nil)
	free()
	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to compile %v: %v", src, log))
	}
	return shader
}

func GetDefaultShaderProgram(vertexShaderSource string, fragmentShaderSource string) uint32 {
	vertexShader := Compile(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := Compile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return shaderProgram
}
func SetUniformMatrix4fv(program uint32, uniform string, matrix [16]float32) {
	projectionUniform := gl.GetUniformLocation(program, gl.Str(uniform+"\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &matrix[0])
}
