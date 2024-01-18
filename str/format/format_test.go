package main

import (
	"fmt"
	"testing"
)

/*
General:

	%v	the value in a default format
	        when printing structs, the plus flag (%+v) adds field names
	%#v	a Go-syntax representation of the value
	%+v: a representation of the fields of the value
	%10v: 10left padding
	%w wrap error: err=fmt.Errorf("wrap %w", err); fmt.Printf("%+v", errors.Unwrap(err))
	%T	a Go-syntax representation of the type of the value
	%%	a literal percent sign; consumes no value
	%q: a single-quoted character literal safely escaped with Go syntax.
	    %+q   Escaped Unicode char(%+q). see: go-str print:
	%T: a Go-syntax representation of the type of the value
	%d: integer in base 10
	%e: scientific notation, e.g. -1.234456e+78
	%E: scientific notation, e.g. -1.234456E+78
	%s: the uninterpreted bytes of the string or slice
	%p: pointer in base 16 notation, with leading 0x

Integer:

	%0.2f: floating-point number
	%b	base 2
	%c	the character represented by the corresponding Unicode code point
	%d	base 10
	%o	base 8
	%x	base 16, with lower-case letters for a-f
	%X	base 16, with upper-case letters for A-F
	%U	Unicode format: U+1234; same as "U+%04X"
*/
func TestFormat(t *testing.T) {
	s := fmt.Sprintf("%.0f", float64(200))
	fmt.Println(s)
}
