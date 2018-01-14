package main

import (
	"os"
	"log"
	"text/template"
	"strings"
	"path/filepath"
)

//+build ignore

type Type struct {
	Import   string // full import line, e.g. `import "github.com/fogleman/gg"`
	Package  string // package, including trailing . (so for builtins, Package=="")
	T        string // name of type in source (including package)
	Name     string // name of type for use in exportable names (capitalized)
	Filename string
}

var EmptyInterfaceType = Type{
	Import:   "",
	Package:  "",
	T:        "interface{}",
	Name:     "Interface",
	Filename: "interface_slice_2d.go",
}

func BuiltinType(t string) Type {
	return Type{
		Import:   "",
		Package:  "",
		T:        t,
		Name:     strings.Title(t),
		Filename: t + "_slice_2d.go",
	}
}

func ExternalType(importPath, pkg, t string) Type {
	return Type{
		Import:   "import \"" + importPath + "\"",
		Package:  pkg + ".",
		T:        pkg + "." + t,
		Name:     strings.Title(t),
		Filename: t + "_slice_2d.go",
	}
}

func main() {
	sliceTemplate := template.Must(template.New("Slice2D").Parse(Slice2DTemplate))

	types := []Type{
		BuiltinType("int"),
		BuiltinType("int8"),
		BuiltinType("int16"),
		BuiltinType("int32"),
		BuiltinType("int64"),
		BuiltinType("uint"),
		BuiltinType("uint8"),
		BuiltinType("uint16"),
		BuiltinType("uint32"),
		BuiltinType("uint64"),
		BuiltinType("string"),
		BuiltinType("float64"),
		BuiltinType("float32"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec2"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec3"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec4"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec2i"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec3i"),
		ExternalType("github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec4i"),
		EmptyInterfaceType,
	}

	for _, t := range types {
		f, err := os.Create(filepath.Join("..", t.Filename))
		//f, err := os.Create(t.Filename)
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
{{.Import}}

// {{.Name}}Slice2D is a 2-dimensional slice of int
type {{.Name}}Slice2D struct {
	W, H int
	data []{{.T}}
}

func New{{.Name}}Slice2D(x, y int) {{.Name}}Slice2D {
	return {{.Name}}Slice2D{
		W:    x,
		H:    y,
		data: make([]{{.T}}, x*y),
	}
}

func (s *{{.Name}}Slice2D) Get(x, y int) {{.T}} {
	return s.data[s.W * y + x]
}

func (s *{{.Name}}Slice2D) Set(x, y int, val {{.T}}) {
	s.data[s.W * y + x] = val
}

func (s *{{.Name}}Slice2D) At(x, y int) *{{.T}} {
	return &s.data[s.W * y + x]
}
`
