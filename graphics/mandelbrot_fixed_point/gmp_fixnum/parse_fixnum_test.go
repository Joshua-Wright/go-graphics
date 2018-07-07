package gmp_fixnum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseFixnum(t *testing.T) {
	assert.Equal(t, "8", ParseFixnum("2.0", 2).String())
	assert.Equal(t, "1289748", ParseFixnum("1.23", 20).String())
	assert.Equal(t, "-569377", ParseFixnum("-0.54,3", 20).String())

	// wolfram alpha: 12.3452345234523678967890912345145234123234523451234532*2^350
	assert.Equal(t, "2831377829361085579558082171789897522961340742409489249405976585072156986425459851534547842962314"+
		"9794821267", ParseFixnum("12.345 23 45,234523678967890912345145234123234523451234532", 350).String())
}
