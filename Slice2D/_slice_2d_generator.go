package main

import (
	"os"
	"log"
	"text/template"
	"strings"
	"fmt"
)

//+build ignore

type Type struct {
	Package  string // package, including trailing . (so for builtins, Package=="")
	T        string // name of type in source (including package)
	Name     string // name of type for use in exportable names (capitalized)
	Filename string
}

var EmptyInterfaceType = Type{
	Package:  "",
	T:        "interface{}",
	Name:     "Interface",
	Filename: "interface_slice_2d.go",
}

func NewBuiltinType(t string) Type {
	return Type{
		Package:  "",
		T:        t,
		Name:     strings.Title(t),
		Filename: fmt.Sprintf("%s_slice_2d.go", t),
	}
}

func main() {
	sliceTemplate := template.Must(template.New("Slice2D").Parse(Slice2DTemplate))

	types := []Type{
		NewBuiltinType("int"),
		NewBuiltinType("int8"),
		NewBuiltinType("int16"),
		NewBuiltinType("int32"),
		NewBuiltinType("int64"),
		NewBuiltinType("uint"),
		NewBuiltinType("uint8"),
		NewBuiltinType("uint16"),
		NewBuiltinType("uint32"),
		NewBuiltinType("uint64"),
		NewBuiltinType("string"),
		NewBuiltinType("float64"),
		NewBuiltinType("float32"),
		EmptyInterfaceType,
	}

	for _, t := range types {
		//f, err := os.Create(filepath.Join("..", t.Filename))
		f, err := os.Create(t.Filename)
		die(err)
		err = sliceTemplate.Execute(f, t)
		die(err)
	}
}
func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Slice2DTemplate = `// generated code, do not edit!
package Slice2D

// {{.Name}}Slice2D is a 2-dimensional slice of int
type {{.Name}}Slice2D struct {
	w, h int
	data []{{.T}}
}

func (s *{{.Name}}Slice2D) Get(x, y int) {{.T}} {
	return s.data[s.w * y + x]
}

func (s *{{.Name}}Slice2D) Set(x, y int, val {{.T}}) {
	s.data[s.w * y + x] = val
}

func (s *{{.Name}}Slice2D) At(x, y int) *{{.T}} {
	return &s.data[s.w * y + x]
}
`
