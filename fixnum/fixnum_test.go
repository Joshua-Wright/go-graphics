package fixnum

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
)

func TestFixnum_SetZero(t *testing.T) {
	f := new(Fixnum).SetZero()
	assert.Equal(t, 0.0, f.ToFloat())
}

func TestFixnum_SetInt(t *testing.T) {
	for i := -20; i < 20; i++ {
		f := new(Fixnum).SetInt(i)
		assert.Equal(t, float64(i), f.ToFloat())
	}
}

func TestFixnum_AddFp(t *testing.T) {
	// integers
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			f := new(Fixnum).SetInt(i)
			g := new(Fixnum).SetInt(j)
			assert.Equal(t, float64(i+j), f.AddFp(f, g).ToFloat())
		}
	}

	rd := rand.New(rand.NewSource(123))

	signs := []int{-1, 1}
	// integers with fractional component
	for i := 0; i < 10; i++ {

		f := new(Fixnum)
		g := new(Fixnum)

		f.sign = signs[rd.Intn(len(signs))]
		f.m[0] = uint32(rd.Int31n(20))
		f.m[1] = rd.Uint32()

		g.sign = signs[rd.Intn(len(signs))]
		g.m[0] = uint32(rd.Int31n(20))
		g.m[1] = rd.Uint32()

		ff := f.ToFloat()
		gf := g.ToFloat()

		expected := ff + gf
		actual := f.AddFp(f, g).ToFloat()
		assert.Equal(t, expected, actual)
	}
}

func TestFixnum_MulFp(t *testing.T) {
	// integers
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			f := new(Fixnum).SetInt(i)
			g := new(Fixnum).SetInt(j)
			assert.Equal(t, float64(i*j), f.MulFp(f, g).ToFloat())
		}
	}

	rd := rand.New(rand.NewSource(123))

	signs := []int{-1, 1}
	// integers with fractional component
	for i := 0; i < 10; i++ {

		f := new(Fixnum)
		g := new(Fixnum)

		f.sign = signs[rd.Intn(len(signs))]
		f.m[0] = uint32(rd.Int31n(20))
		f.m[1] = rd.Uint32()

		g.sign = signs[rd.Intn(len(signs))]
		g.m[0] = uint32(rd.Int31n(20))
		g.m[1] = rd.Uint32()

		ff := f.ToFloat()
		gf := g.ToFloat()

		expected := ff * gf
		actual := f.MulFp(f, g).ToFloat()
		assert.Equal(t, expected, actual)
	}
}
