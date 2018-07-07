package main

import (
	"math/big"
	g "github.com/joshua-wright/go-graphics/graphics"
	"fmt"
)

/*
decimalDigits = 100
decimal[x_] := DecimalForm[N[x, decimalDigits], decimalDigits]
x := 0.000354852939810431279017857142857142857142857142857142857142857\
14285714285714285714285714285714285714285714285714285714285714285714
steps := 5120
dx := 2/35840000000000000
delta := dx/steps
decimal[delta]
decimal[dx + x]
 */
func main() {
	num := "0.00035485293981043127901785714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714285714"
	steps := 5120
	delta := "0.00000000000000000001089913504464285714285714285714285714285714285714285714285714285714285714285714285714285714285714286"
	final := "0.0003548529398104870825892857142857142857142857142857142857142857142857142857142857142857142857142857143"
	finalNeg := "-0.0003548529398104870825892857142857142857142857142857142857142857142857142857142857142857142857142857143"

	bits := uint(96)

	{ // big float
		x, _, err := big.ParseFloat(num, 10, bits, big.ToNearestEven)
		g.Die(err)
		delta_, _, err := big.ParseFloat(delta, 10, bits, big.ToNearestEven)
		g.Die(err)
		final_, _, err := big.ParseFloat(final, 10, bits, big.ToNearestEven)
		g.Die(err)

		for i := 0; i < steps; i++ {
			x.Add(x, delta_)
		}
		x.Sub(x, final_)
		fmt.Println(x.Text('f', int(bits)))
	}

	{ // fixnum
		x, err := naive_fixnum.FromString(num)
		g.Die(err)
		delta_, err := naive_fixnum.FromString(delta)
		g.Die(err)
		//final_, err := fixnum.FromString(final)
		//g.Die(err)
		finalNeg_, err := naive_fixnum.FromString(finalNeg)
		g.Die(err)

		for i := 0; i < steps; i++ {
			x.Add(x, delta_)
		}
		x.Add(x, finalNeg_)
		fmt.Println(x)
	}
}
