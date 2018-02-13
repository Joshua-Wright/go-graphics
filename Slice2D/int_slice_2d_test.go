package Slice2D

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntSlice2D(t *testing.T) {
	w := 10
	h := 20
	slice := NewIntSlice2D(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			slice.Set(i, j, i*j*j)
		}
	}
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			assert.Equal(t, slice.Get(i, j), i*j*j)
			assert.Equal(t, slice.At(i, j), &slice.Data[j*w+i])
		}
	}
}
