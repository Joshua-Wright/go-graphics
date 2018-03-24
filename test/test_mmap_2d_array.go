package main

import (
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
)

func main() {
	width := int64(1024 * 3)
	height := int64(1024 * 5)

	func() {
		arr, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "test_float64_2d_array")
		g.Die(err)

		for j := int64(0); j < height; j++ {
			for i := int64(0); i < width; i++ {
				arr.Set(i, j, float64(i*j)+float64(i)*0.5)
			}
		}
		arr.Close()
	}()

	func() {
		arr, err := memory_mapped.OpenOrCreateMmap2dArrayFloat64(width, height, "test_float64_2d_array")
		g.Die(err)

		for j := int64(0); j < height; j++ {
			for i := int64(0); i < width; i++ {
				if arr.Get(i, j) != float64(i*j)+float64(i)*0.5 {
					fmt.Println("failed at", i, arr.Get(i, j), float64(i*j)+float64(i)*0.5)
				}
			}
		}
		arr.Close()
	}()

}
