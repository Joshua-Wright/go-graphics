package main

import (
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
)

func main() {
	length := int64(1024*10)

	func() {
		arr, err := memory_mapped.OpenOrCreateMmapArrayFloat64(length, "test_float64_array")
		g.Die(err)

		for i := int64(0); i < length; i++ {
			arr.Set(i, float64(i*i))
		}
		arr.Close()
	}()

	func() {
		arr, err := memory_mapped.OpenOrCreateMmapArrayFloat64(length, "test_float64_array")
		g.Die(err)

		for i := int64(0); i < length; i++ {
			if arr.Get(i) != float64(i*i) {
				fmt.Println("failed at", i, arr.Get(i), float64(i*i))
			}
		}
		arr.Close()
	}()

}
