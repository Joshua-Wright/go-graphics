package main

import "fmt"
import m "github.com/joshua-wright/image_stuff/go"

func main() {
	v := m.Vec3{1, 2, 3}
	fmt.Print(v)
	r := m.Ray{m.Vec3{0, 0, 0}, m.Vec3{1, 1, 1}}
	s := m.Sphere{
		m.Vec3{3, 3, 3},
		0.5,
	}
	fmt.Println(s.Intersect(&r))
	fmt.Println(m.Vec2{1, 2})
	fmt.Println(m.Vec3{1, 2, 3})
	fmt.Println(m.Vec4{1, 2, 3, 4})
}
