package main

import (
	"github.com/joshua-wright/go-graphics/graphics/memory_mapped"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
	"github.com/joshua-wright/go-graphics/parallel"
	"runtime"
)

func main() {
	length := int64(1024 * 10)

	modm := int64(1234)

	procs := runtime.GOMAXPROCS(-1)

	func() {
		arr, err := memory_mapped.OpenOrCreateAtomicBitset(length, "test_atomic_bitset")
		g.Die(err)

		parallel.ParallelFor(0, procs, func(_ int) {
			for i := int64(procs); i < length; i += int64(procs) {
				if i%modm == 0 {
					arr.Set(i)
				}
			}
		})

		arr.Close()
	}()

	func() {
		arr, err := memory_mapped.OpenOrCreateAtomicBitset(length, "test_atomic_bitset")
		g.Die(err)

		parallel.ParallelFor(0, procs, func(_ int) {
			for i := int64(procs); i < length; i += int64(procs) {
				if arr.Test(i) != (i%modm == 0) {
					fmt.Println("failed at", i, arr.Test(i), (i%modm == 0))
				}
			}
		})

		arr.Close()
	}()

}
