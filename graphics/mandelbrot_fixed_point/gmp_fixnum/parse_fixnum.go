package gmp_fixnum

import (
	"github.com/ncw/gmp"
	"strings"
	"regexp"
	"github.com/pkg/errors"
)

var digitPastDecimalRegex, regexErr = regexp.Compile(`\.\d*`)

func ParseFixnum(decimal string, base_power_2 uint) *gmp.Int {
	if regexErr != nil {
		panic(regexErr)
	}
	if (strings.Count(decimal, ".") != 1) {
		panic("bad number format")
	}
	// remove whitespace things
	decimal = strings.Replace(decimal, " ", "", -1)
	decimal = strings.Replace(decimal, ",", "", -1)

	numerator := strings.Replace(decimal, ".", "", -1)

	denominator_power_10 := len(digitPastDecimalRegex.FindString(decimal)) - 1
	denominator := gmp.NewInt(10)
	denominator.Exp(denominator, gmp.NewInt(int64(denominator_power_10)), nil)

	val := new(gmp.Int)
	val.SetString(numerator, 10)
	val.Lsh(val, base_power_2)
	val.Div(val, denominator)

	return val
}

// TODO: combine with above
func ParseFixnumSafe(decimal string, base_power_2 uint) (*gmp.Int, error) {
	if regexErr != nil {
		panic(regexErr)
	}
	if (strings.Count(decimal, ".") != 1) {
		return nil, errors.New("bad number format: " + decimal)
	}
	// remove whitespace things
	decimal = strings.Replace(decimal, " ", "", -1)
	decimal = strings.Replace(decimal, ",", "", -1)

	numerator := strings.Replace(decimal, ".", "", -1)

	denominator_power_10 := len(digitPastDecimalRegex.FindString(decimal)) - 1
	denominator := gmp.NewInt(10)
	denominator.Exp(denominator, gmp.NewInt(int64(denominator_power_10)), nil)

	val := new(gmp.Int)
	val.SetString(numerator, 10)
	val.Lsh(val, base_power_2)
	val.Div(val, denominator)

	return val, nil
}
