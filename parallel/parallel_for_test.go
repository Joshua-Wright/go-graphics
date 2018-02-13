package parallel

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/mkideal/pkg/math/random"
)

func TestParallelFor(t *testing.T) {
	P := 10
	ints := make([]int, P)
	ParallelFor(0, P, func(i int) {
		ints[i] = i
	})
	ParallelFor(0, P, func(i int) {
		assert.Equal(t, ints[i], i)
	})
	for i := 0; i < P; i++ {
		assert.Equal(t, ints[i], i)
	}
}

func TestParallelFuncs(t *testing.T) {
	a := 0
	b := 0
	c:= 0
	ParallelFuncs(
		func(){a = 1},
		func(){b = 2},
		func(){c = 3},
	)
	assert.Equal(t, a, 1)
	assert.Equal(t, b, 2)
	assert.Equal(t, c, 3)
}

func TestParallelForAdaptive(t *testing.T) {
	P := 1000
	ints := make([]int, P)
	elapsed := ParallelForAdaptive(0, P, func(startInclusive, endExclusive int) {
		for i := startInclusive; i < endExclusive; i++ {
			ns := 1000000+random.Intn(10000, random.DefaultSource)
			time.Sleep(time.Duration(ns)*time.Nanosecond)
			ints[i] = i
		}
	})
	println(elapsed.String())
	for i := 0; i < P; i++ {
		assert.Equal(t, ints[i], i)
	}
}
