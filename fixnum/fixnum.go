package fixnum

// thanks to: http://www.bealto.com/mp-mandelbrot_fp-reals.html

const fpWords = 4

// lower 32-bits set
const wordMask uint64 = 0x00000000FFFFFFFF

type Fixnum struct {
	sign int
	m    [fpWords]uint32
}

func (z *Fixnum) checkZero() {
	if z.sign == 0 {
		return
	}
	for i := 0; i < fpWords; i++ {
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
	for i := int(fpWords - 1); i >= 0; i-- {
		c += uint64(z.m[i]) + uint64(g.m[i])
		z.m[i] = uint32(c & wordMask)
		c >>= 32
	}
	return z
}

func (z *Fixnum) subWords(x *Fixnum) *Fixnum {
	var c uint64 = 0
	for i := int(fpWords - 1); i >= 0; i-- {
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
	for i := 0; i < fpWords; i++ {
		if z.m[i] > x.m[i] {
			return 1
		};
		if z.m[i] < x.m[i] {
			return -1
		}
	}
	return 0;
}

func (z *Fixnum) AddFp(x, y *Fixnum) *Fixnum {
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

func (z *Fixnum) MulFp(x, y *Fixnum) *Fixnum {
	if x.sign == 0 || y.sign == 0 {
		z.SetZero()
		return z
	}

	z.sign = x.sign * y.sign

	// multiply (trivial way)
	var aux [fpWords]uint64
	for i := 0; i < fpWords; i++ {
		for j := 0; j < fpWords; j++ {
			k := i + j
			if k > fpWords {
				continue
			}
			u1 := uint64(x.m[i]) * uint64(y.m[j]);
			u0 := u1 & wordMask; // lower 32 bits, index K
			u1 >>= 32;           // higher 32 bits, index K-1
			if k < fpWords {
				aux[k] += u0
			}
			if k > 0 {
				aux[k-1] += u1
			}
		}
	}

	// propagate carry
	var c uint64
	for i := int(fpWords - 1); i >= 0; i-- {
		c += aux[i];
		z.m[i] = uint32(c & wordMask);
		c >>= 32;
	}

	return z
}

// approximation
const c1 = 1.0 / 4294967296.0           // 1 / 2^32
const c2 = 1.0 / 18446744073709551616.0 // 1 / 2^64
func (z *Fixnum) ToFloat() float64 {
	if z.sign == 0 {
		return 0.0
	}
	return float64(z.sign) * (float64(z.m[0]) + c1*float64(z.m[1]) + c2*float64(z.m[2]))
}
