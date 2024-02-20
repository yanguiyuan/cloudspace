// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reflection_tpl

// File .
var File = `// Code generated by thriftgo ({{Version}}). DO NOT EDIT.
{{InsertionPoint "bof"}}

package {{.FilePackage}}

import (
	{{InsertionPoint "imports"}}
	{{define "Imports"}}
	{{end}}
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

{{$IDLName := .IDLName}}
{{$IDLPath := .AST.Filename}}
{{$FilePackage := .FilePackage}}

// IDL Name: {{$IDLName}}
// IDL Path: {{$IDLPath}}

var file_{{$IDLName}}_thrift_go_types = []interface{}{
	{{- range $index, $element := .Structs}}
	(*{{.GoName}})(nil),	// Struct {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Unions}}
	(*{{.GoName}})(nil),	// Union {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Exceptions}}
	(*{{.GoName}})(nil),	// Exception {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Enums}}
	(*{{.GoName}})(nil),	// Enum {{$index}}: {{$FilePackage}}.{{.Name}}
	{{- end}}
	{{- range $index, $element := .Typedefs}}
	(*{{.GoName}})(nil),	// Enum {{$index}}: {{$FilePackage}}.{{.Alias}}
	{{- end}}
}
var file_{{$IDLName}}_thrift *thrift_reflection.FileDescriptor
var file_idl_{{$IDLName}}_rawDesc = {{.MarshalDescriptor}}

func init() {
	if file_{{$IDLName}}_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes: file_idl_{{$IDLName}}_rawDesc,
		GoTypes:file_{{$IDLName}}_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_{{$IDLName}}_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorFor{{ToCamel $IDLName}}() *thrift_reflection.FileDescriptor{
	return file_{{$IDLName}}_thrift
}

{{- range .Structs}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetStructDescriptor("{{.Name}}")
}

func (p *{{.GoName}}) GetTypeDescriptor() *thrift_reflection.TypeDescriptor{
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_{{$IDLName}}_thrift.Filepath
	ret.Name = "{{.Name}}"
	return ret
}
{{- end}}
{{- range .Enums}}
func (p {{.GoName}}) GetDescriptor() *thrift_reflection.EnumDescriptor{
	return file_{{$IDLName}}_thrift.GetEnumDescriptor("{{.Name}}")
}

func (p *{{.GoName}}) GetTypeDescriptor() *thrift_reflection.TypeDescriptor{
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_{{$IDLName}}_thrift.Filepath
	ret.Name = "{{.Name}}"
	return ret
}
{{- end}}
{{- range .Unions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetUnionDescriptor("{{.Name}}")
}

func (p *{{.GoName}}) GetTypeDescriptor() *thrift_reflection.TypeDescriptor{
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_{{$IDLName}}_thrift.Filepath
	ret.Name = "{{.Name}}"
	return ret
}
{{- end}}
{{- range .Exceptions}}
func (p *{{.GoName}}) GetDescriptor() *thrift_reflection.StructDescriptor{
	return file_{{$IDLName}}_thrift.GetExceptionDescriptor("{{.Name}}")
}

func (p *{{.GoName}}) GetTypeDescriptor() *thrift_reflection.TypeDescriptor{
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_{{$IDLName}}_thrift.Filepath
	ret.Name = "{{.Name}}"
	return ret
}
{{- end}}


{{- InsertionPoint "eof"}}
`
