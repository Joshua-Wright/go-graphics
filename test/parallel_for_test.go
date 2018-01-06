package main
import (
	m "github.com/joshua-wright/go-graphics/graphics"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParallelFor(t *testing.T) {
	P := 10
	ints := make([]int, P)
	m.ParallelFor(0, P, func(i int) {
		ints[i] = i
	})
	m.ParallelFor(0, P, func(i int) {
		assert.Equal(t, ints[i], i)
	})
	for i := 0; i < P; i++ {
		assert.Equal(t, ints[i], i)
	}
}
