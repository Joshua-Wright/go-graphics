package main

// +build ignore

import (
	"os"
	"log"
	"text/template"
	"io"
	"fmt"
)

func main() {
	f, err := os.Create("generated_vectors.go")
	die(err)
	fmt.Fprintln(f, `package graphics

import "fmt"`)

	v2 := Type{
		F:      f,
		T:      "Vec2",
		Fields: []string{"X", "Y"},
	}
	v3 := Type{
		F:      f,
		T:      "Vec3",
		Fields: []string{"X", "Y", "Z"},
	}
	v4 := Type{
		F:      f,
		T:      "Vec4",
		Fields: []string{"X", "Y", "Z", "W"},
	}

	for _, t := range []Type{v2, v3, v4} {
		t.WriteType()
		t.Clone()

		t.BinaryVectorOp("+", "Add")
		t.BinaryVectorOp("-", "Sub")
		t.BinaryVectorOp("*", "Mul")
		t.BinaryVectorOp("/", "Div")

		t.BinaryScalarOp("+", "Add")
		t.BinaryScalarOp("-", "Sub")
		t.BinaryScalarOp("*", "Mul")
		t.BinaryScalarOp("/", "Div")

		t.Mag()
		t.Dot()
		t.Constants()
		t.MakeString()
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Type struct {
	F      io.Writer
	T      string
	Fields []string
}

func (t *Type) WriteType() {
	// TODO don't hard code this
	err := template.Must(template.New("").Parse(`
//go:generate go run ../codegen/generate_vectors.go
type {{.T}} struct {
{{- range .Fields }}
	{{ . }} Float
{{- end }}
}
`)).Execute(t.F, t)
	die(err)
}

func (t *Type) BinaryVectorOp(op, name string) {
	// TODO generate tests
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) {{.OpName}}V(u {{.T}}) {{.T}} {
	return {{.T}} { {{ $Op := .Op }}
		{{- range .Fields }}
		{{ . }}: v.{{ . }} {{$Op}} u.{{ . }},
		{{- end }}
	}
}
//func (v *{{.T}}) {{.OpName}}Vi(u *{{.T}}) *{{.T}} {
//	{{ $Op := .Op }} {{- range .Fields }}
//	v.{{ . }} {{$Op}}= u.{{ . }}
//	{{- end }}
//	return v
//}
`)).Execute(
		t.F, struct {
			Type
			Op     string
			OpName string
		}{
			Type:   *t,
			Op:     op,
			OpName: name,
		})
	die(err)
}

func (t *Type) BinaryScalarOp(op, name string) {
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) {{.OpName}}S(s Float) {{.T}} {
	return {{.T}} { {{ $Op := .Op }}
		{{- range .Fields }}
		{{ . }}: v.{{ . }} {{$Op}} s,
		{{- end }}
	}
}
//func (v *{{.T}}) {{.OpName}}Si(s Float) *{{.T}} {
//	{{ $Op := .Op }} {{- range .Fields }}
//	v.{{ . }} {{$Op}}= s
//	{{- end }}
//	return v
//}
`)).Execute(
		t.F, struct {
			Type
			Op     string
			OpName string
		}{
			Type:   *t,
			Op:     op,
			OpName: name,
		})
	die(err)
}

func (t *Type) Clone() {
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) Clone() *{{.T}} {
	return &{{.T}} {
		{{- range .Fields }}
		{{ .  }}: v.{{ . }},
		{{- end }}
	}
}
`)).Execute(
		t.F, t)
	die(err)
}

func (t *Type) Dot() {
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) Dot(u {{.T}}) Float {
	return {{- range $i, $f := .Fields }} {{if $i}}+{{end}} v.{{.}}*v.{{.}}{{- end }}
}
`)).Execute(
		t.F, t)
	die(err)
}

func (t *Type) Mag() {
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) Mag2() Float {
	return {{- range $i, $f := .Fields }} {{if $i}}+{{end}} v.{{.}}*v.{{.}}{{- end }}
}
func (v {{.T}}) Mag() Float {
	return Sqrt({{- range $i, $f := .Fields }}{{if $i}}+{{end}} v.{{.}}*v.{{.}}{{- end }})
}
func (v {{.T}}) UnitV() {{.T}} {
	l := Float(1.0)/Sqrt({{- range $i, $f := .Fields }}{{if $i}}+{{end}} v.{{.}}*v.{{.}}{{- end }})
	return {{.T}} {
		{{- range .Fields }}
		{{ . }}: v.{{ . }}*l,
		{{- end }}
	}
}
//func (v *{{.T}}) UnitVi() *{{.T}} {
//	l := Float(1.0)/Sqrt({{- range $i, $f := .Fields }}{{if $i}}+{{end}} v.{{.}}*v.{{.}}{{- end }})
//	{{- range .Fields }}
//	v.{{.}} *= l
//	{{- end }}
//	return v
//}
`)).Execute(
		t.F, t)
	die(err)
}

func (t *Type) Constants() {
	err := template.Must(template.New("").Parse(`
var {{.T}}Zero = {{.T}} {
	{{- range $i, $f := .Fields }}
		{{.}}: 0,
	{{- end }}
}
var {{.T}}One = {{.T}} {
	{{- range $i, $f := .Fields }}
		{{.}}: 1,
	{{- end }}
}
{{ $T := .T }}
{{ $Fields := .Fields }}
{{range $i, $f := .Fields }}
var {{$T}}{{.}} = {{$T}} {
	{{- range $i, $f2 := $Fields }}
	{{.}}: {{if eq $f $f2 }} 1, {{ else }} 0, {{ end }}
	{{- end }}
}
{{ end }}
`)).Execute(
		t.F, t)
	die(err)
}

func (t *Type) MakeString() {
	err := template.Must(template.New("").Parse(`
func (v {{.T}}) String() string {
	return fmt.Sprint("[", {{- range $i, $f := .Fields }} {{- if $i}}, ",",{{end}} v.{{.}}{{- end }}, "]")
}
`)).Execute(
		t.F, t)
	die(err)
}
