package naive_fixnum_62

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"math/big"
	"math"
	"fmt"
)

func TestFixnum_SetZero(t *testing.T) {
	f := new(Fixnum).SetZero()
	assert.Equal(t, 0.0, f.Float64())
}

func TestFixnum_SetInt(t *testing.T) {
	for i := -20; i < 20; i++ {
		f := new(Fixnum).SetInt(i)
		assert.Equal(t, float64(i), f.Float64())
	}
}

func TestFixnum_SubFp(t *testing.T) {
	// integers
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			f := new(Fixnum).SetInt(i)
			g := new(Fixnum).SetInt(j)
			assert.Equal(t, float64(i-j), f.Sub(f, g).Float64())
		}
	}

	rd := rand.New(rand.NewSource(123))

	signs := []int32{-1, 1}
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

		ff := f.Float64()
		gf := g.Float64()

		expected := ff - gf
		actual := f.Sub(f, g).Float64()
		assert.Equal(t, expected, actual)
	}
}

func TestFixnum_AddFp(t *testing.T) {
	// integers
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			f := new(Fixnum).SetInt(i)
			g := new(Fixnum).SetInt(j)
			assert.Equal(t, float64(i+j), f.Add(f, g).Float64())
		}
	}

	rd := rand.New(rand.NewSource(123))

	signs := []int32{-1, 1}
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

		ff := f.Float64()
		gf := g.Float64()

		expected := ff + gf
		actual := f.Add(f, g).Float64()
		assert.Equal(t, expected, actual)
	}
}

func TestFixnum_MulFp(t *testing.T) {
	// integers
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			f := new(Fixnum).SetInt(i)
			g := new(Fixnum).SetInt(j)
			assert.Equal(t, float64(i*j), f.Mul(f, g).Float64())
		}
	}

	rd := rand.New(rand.NewSource(123))

	signs := []int32{-1, 1}
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

		ff := f.Float64()
		gf := g.Float64()

		expected := ff * gf
		actual := f.Mul(f, g).Float64()
		assert.Equal(t, expected, actual)
	}
}

func TestFromBigFloat(t *testing.T) {
	rd := rand.New(rand.NewSource(123))
	for i := 0; i < 100; i++ {
		expected := rd.Float64()
		bf := big.NewFloat(expected)
		f := FromBigFloat(bf)
		assert.Equal(t, expected, f.Float64())
	}

	assert.Panics(t, func() {
		expected := math.Pow(2, 70)
		bf := big.NewFloat(expected)
		f := FromBigFloat(bf)
		assert.Equal(t, expected, f.Float64())
	})
}

func TestFromString(t *testing.T) {
	rd := rand.New(rand.NewSource(123))
	for i := 0; i < 100; i++ {
		expected := rd.Float64()
		f, err := FromString(fmt.Sprint(expected))
		assert.Nil(t, err)
		assert.Equal(t, expected, f.Float64())
	}
}
