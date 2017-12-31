package main

import (
	m "github.com/joshua-wright/image_stuff/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonSquareMakeBounds(t *testing.T) {
	width := 2000
	height := 1000
	bounds := m.BoundsForResolution(0, 0, 2, width, height)
	xmin := bounds[0]
	xmax := bounds[1]
	ymin := bounds[2]
	ymax := bounds[3]
	assert.Equal(t, float64(-1), xmin)
	assert.Equal(t, float64(1), xmax)
	assert.Equal(t, -0.5, ymin)
	assert.Equal(t, 0.5, ymax)
}
func TestNonSquareWindowTransform(t *testing.T) {
	width := 2000
	height := 1000
	bounds := m.BoundsForResolution(0, 0, 2, width, height)

	corners := []m.Vec2{
		{-1, 0.5},
		{-1, -0.5},
		{1, 0.5},
		{1, -0.5},
	}
	corners_window := []m.Vec2{
		{0, 0},
		{0, float64(height - 1)},
		{float64(width - 1), 0},
		{float64(width - 1), float64(height - 1)},
	}

	for i := 0; i < len(corners); i++ {
		xi, yi := m.WindowTransformPoint(width, height, corners[i], bounds)
		assert.Equal(t, int(corners_window[i].X), xi)
		assert.Equal(t, int(corners_window[i].Y), yi)
	}

}
