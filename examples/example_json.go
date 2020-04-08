package examples

import (
	"strings"
	"unicode"

	. "github.com/ihrk/goparsec"
)

func jsonParser() Combinator {
	digits := Some(RuneFunc(unicode.IsDigit))
	trimSpaces := Many(RuneFunc(unicode.IsSpace))
	comma := And(Rune(','), trimSpaces)
	quot := Rune('"')
	jNum := And(
		Opt(Rune('-')),
		digits,
		Opt(And(
			Rune('.'),
			digits),
		))
	jStr := And(
		quot,
		Many(
			Or(
				Some(RuneFunc(func(r rune) bool { return !strings.ContainsRune("\"\b\f\n\r\t\\", r) })),
				And(
					Rune('\\'),
					RuneFunc(func(r rune) bool { return strings.ContainsRune(`"\/bfnrt`, r) }),
				))),
		quot,
	)
	jBool := Or(Str("true"), Str("false"))
	jNull := Str("null")
	jObj, jArr := And(), And()
	jVal := Or(jNull, jBool, jNum, jStr, jArr, jObj)
	*jObj = *And(
		Rune('{'),
		trimSpaces,
		SepBy(
			And(jStr, trimSpaces, Rune(':'), trimSpaces, jVal, trimSpaces),
			comma,
		),
		Rune('}'),
	)
	*jArr = *And(
		Rune('['),
		trimSpaces,
		SepBy(
			And(
				jVal,
				trimSpaces,
			),
			comma,
		),
		Rune(']'),
	)
	return And(trimSpaces, jVal, trimSpaces, EOF())
}
