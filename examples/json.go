package examples

import (
	"unicode"

	. "github.com/ihrk/goparsec"
)

func jsonParser() ParserFunc {
	digits := TrimFunc1(unicode.IsDigit)
	comma := Sequence(Rune(','), TrimSpace)
	quot := Rune('"')
	jNum := Sequence(
		Optional(Rune('-')),
		digits,
		Optional(Sequence(
			Rune('.'),
			digits),
		))
	jStr := Sequence(
		quot,
		Many(
			Choice(
				TrimFunc1(
					func(r rune) bool {
						return r != '"' && r != '\b' &&
							r != '\f' && r != '\n' &&
							r != '\r' && r != '\t' &&
							r != '\\'
					}),
				Sequence(
					Rune('\\'),
					Func(func(r rune) bool {
						return r == '"' || r == '\\' ||
							r == '/' || r == 'b' ||
							r == 'f' || r == 'n' ||
							r == 'r' || r == 't'
					}),
				),
			)),
		quot,
	)
	jBool := Choice(Str("true"), Str("false"))
	jNull := Str("null")
	jObj, jArr := NewBox(), NewBox()
	jVal := Choice(jNull, jBool, jNum, jStr, jArr.Unwrap(), jObj.Unwrap())
	jObj.Wrap(Sequence(
		Rune('{'),
		TrimSpace,
		SepBy(
			Sequence(jStr, TrimSpace, Rune(':'), TrimSpace, jVal, TrimSpace),
			comma,
		),
		Rune('}'),
	))
	jArr.Wrap(Sequence(
		Rune('['),
		TrimSpace,
		SepBy(
			Sequence(
				jVal,
				TrimSpace,
			),
			comma,
		),
		Rune(']'),
	))
	return Sequence(TrimSpace, jVal, TrimSpace, EOF)
}
