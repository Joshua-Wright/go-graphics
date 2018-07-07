package naive_fixnum

import (
	"math/big"
)

//////////////////////////////////////////////////////////////////////
// thanks to: http://www.bealto.com/mp-mandelbrot_fp-reals.html

// lower 32-bits set
const wordMask uint64 = 0x00000000FFFFFFFF

type Fixnum struct {
	sign int32
	m    [FpWords]uint32
}

func (z *Fixnum) checkZero() {
	if z.sign == 0 {
		return
	}
	for i := 0; i < FpWords; i++ {
		if z.m[i] != 0 {
			return
		}
	}
	z.sign = 0
}

func (z *Fixnum) Set(x *Fixnum) *Fixnum {
	if z == x {
		return z
	}
	*z = *x
	return z
}

func (z *Fixnum) SetZero() *Fixnum {
	*z = Fixnum{}
	return z
}

func (z *Fixnum) SetInt(x int) *Fixnum {
	*z = Fixnum{}
	if x > 0 {
		z.m[0] = uint32(x)
		z.sign = 1
	} else if x < 0 {
		z.m[0] = uint32(-x)
		z.sign = -1
	}
	return z
}

func (z *Fixnum) addWords(g *Fixnum) *Fixnum {
	var c uint64 = 0
	for i := int(FpWords - 1); i >= 0; i-- {
		c += uint64(z.m[i]) + uint64(g.m[i])
		z.m[i] = uint32(c & wordMask)
		c >>= 32
	}
	return z
}

func (z *Fixnum) subWords(x *Fixnum) *Fixnum {
	var c uint64 = 0
	for i := int(FpWords - 1); i >= 0; i-- {
		y := uint64(z.m[i]) - uint64(x.m[i]) - c
		z.m[i] = uint32(y & wordMask)
		if y >= 0x100000000 {
			c = 1
		} else {
			c = 0
		}
	}
	return z
}

func (z *Fixnum) cmpWords(x *Fixnum) int {
	for i := 0; i < FpWords; i++ {
		if z.m[i] > x.m[i] {
			return 1
		};
		if z.m[i] < x.m[i] {
			return -1
		}
	}
	return 0;
}

func (z *Fixnum) Neg() *Fixnum {
	z.sign = -z.sign
	return z
}

func (z *Fixnum) Sub(x, y *Fixnum) *Fixnum {
	y2 := *y
	y2.Neg()
	z.Add(x, &y2)
	return z
}

func (z *Fixnum) Add(x, y *Fixnum) *Fixnum {
	z.Set(x)
	if y.sign == 0 {
		return z
	}
	if z.sign == 0 {
		z.Set(y)
		return z
	}

	if z.sign == y.sign {
		z.addWords(y)
		return z
	}

	// opposite signs, must subtract
	if z.cmpWords(y) >= 0 {
		z.subWords(y)
	} else {
		y2 := *z
		z.Set(y)
		z.subWords(&y2)
	}
	z.checkZero()
	return z
}

func (z *Fixnum) Mul(x, y *Fixnum) *Fixnum {
	if x.sign == 0 || y.sign == 0 {
		z.SetZero()
		return z
	}

	z.sign = x.sign * y.sign

	// multiply (trivial way)
	var aux [FpWords]uint64
	for i := 0; i < FpWords; i++ {
		for j := 0; j < FpWords; j++ {
			k := i + j
			if k > FpWords {
				continue
			}
			u1 := uint64(x.m[i]) * uint64(y.m[j]);
			u0 := u1 & wordMask; // lower 32 bits, index K
			u1 >>= 32;           // higher 32 bits, index K-1
			if k < FpWords {
				aux[k] += u0
			}
			if k > 0 {
				aux[k-1] += u1
			}
		}
	}

	// propagate carry
	var c uint64
	for i := int(FpWords - 1); i >= 0; i-- {
		c += aux[i];
		z.m[i] = uint32(c & wordMask);
		c >>= 32;
	}

	return z
}

// approximation
const c1 = 1.0 / 4294967296.0           // 1 / 2^32
const c2 = 1.0 / 18446744073709551616.0 // 1 / 2^64
func (z *Fixnum) Float64() float64 {
	if z.sign == 0 {
		return 0.0
	}
	return float64(z.sign) * (float64(z.m[0]) + c1*float64(z.m[1]) + c2*float64(z.m[2]))
}

func FromBigFloat(bf *big.Float) *Fixnum {
	return new(Fixnum).FromBigFloat(bf)
}

func (f *Fixnum) FromBigFloat(bf *big.Float) *Fixnum {
	f.sign = int32(bf.Sign())
	if f.sign == 0 {
		return f
	}

	if f.sign == -1 {
		bf.Neg(bf)
	}

	for i := 0; i < FpWords; i++ {
		b, acc := bf.Uint64()
		if acc == big.Above || b > wordMask {
			panic("failed to parse float, maybe number is too big?")
		}
		f.m[i] = uint32(b & wordMask)
		bf.Sub(bf, new(big.Float).SetInt64(int64(b)))
		mant := new(big.Float)
		exp := bf.MantExp(mant)
		bf.SetMantExp(mant, exp+32)
	}

	return f
}

func FromString(str string) (*Fixnum, error) {
	return new(Fixnum).FromString(str)
}

func (f *Fixnum) FromString(str string) (*Fixnum, error) {
	bf, _, err := big.ParseFloat(str, 10, FpWords*32*2, big.ToNearestEven)
	if err != nil {
		return nil, err
	}
	return f.FromBigFloat(bf), nil
}

func FromInt32(i int32) *Fixnum {
	return new(Fixnum).FromInt32(i)
}

func (f *Fixnum) FromInt32(i int32) *Fixnum {
	if i < 0 {
		f.sign = -1
		f.m[0] = uint32(-i)
	} else if i == 0 {
		f.sign = 0
	} else {
		f.sign = 1
		f.m[0] = uint32(i)
	}
	return f
}
