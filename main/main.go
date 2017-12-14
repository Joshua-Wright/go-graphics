package main

import "fmt"
import m ".."

func main() {
	v := m.Vec3{1, 2, 3}
	fmt.Print(v)
	r := m.Ray{m.Vec3{0, 0, 0}, m.Vec3{1, 1, 1}}
	s := m.Sphere{
		m.Vec3{3,3,3},
		0.5,
	}
	fmt.Print(s.Intersect(&r))

}
