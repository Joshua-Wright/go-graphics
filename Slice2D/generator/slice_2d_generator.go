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
	Package  string // package to declare
	T        string // name of type in source (including package)
	Name     string // name of type for use in exportable names (capitalized)
	Filename string
}

var EmptyInterfaceType = Type{
	Import:   "",
	Package:  "Slice2D",
	T:        "interface{}",
	Name:     "Interface",
	Filename: "interface_slice_2d.go",
}

func BuiltinType(t string) Type {
	return Type{
		Import:   "",
		Package:  "Slice2D",
		T:        t,
		Name:     strings.Title(t),
		Filename: t + "_slice_2d.go",
	}
}

func ExternalType(codePackage, importPath, typePackage, t string) Type {
	return Type{
		Import:   "import \"" + importPath + "\"",
		Package:  codePackage,
		T:        typePackage + "." + t,
		Name:     strings.Title(t),
		Filename: t + "_slice_2d.go",
	}
}

func PackageLocalType(pkg, t string) Type {
	return Type{
		Import:   "",
		Package:  pkg,
		T:        t,
		Name:     strings.Title(t),
		Filename: t + "_slice_2d.go",
	}
}

func main() {
	sliceTemplate := template.Must(template.New("Slice2D").Parse(Slice2DTemplate))

	if len(os.Args) > 1 {
		var t Type
		if os.Args[1] == "ExternalType" {
			t = ExternalType(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		} else if os.Args[1] == "PackageLocalType" {
			t = PackageLocalType(os.Args[2], os.Args[3])
		} else {
			log.Fatal("bad args:", os.Args)
		}
		f, err := os.Create(t.Filename)
		die(err)
		err = sliceTemplate.Execute(f, t)
		die(err)
	} else {
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
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec2"),
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec3"),
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec4"),
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec2i"),
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec3i"),
			ExternalType("Slice2D", "github.com/joshua-wright/go-graphics/graphics", "graphics", "Vec4i"),
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
}
func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Slice2DTemplate = `// generated code, do not edit!
package {{.Package}}
{{.Import}}

// {{.Name}}Slice2D is a 2-dimensional slice of int
type {{.Name}}Slice2D struct {
	W, H int
	Data []{{.T}}
}

func New{{.Name}}Slice2D(x, y int) {{.Name}}Slice2D {
	return {{.Name}}Slice2D{
		W:    x,
		H:    y,
		Data: make([]{{.T}}, x*y),
	}
}

func (s *{{.Name}}Slice2D) Get(x, y int) {{.T}} {
	return s.Data[s.W * y + x]
}

func (s *{{.Name}}Slice2D) Set(x, y int, val {{.T}}) {
	s.Data[s.W * y + x] = val
}

func (s *{{.Name}}Slice2D) At(x, y int) *{{.T}} {
	return &s.Data[s.W * y + x]
}
`
