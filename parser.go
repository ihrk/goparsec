package goparsec

import (
	"math"
	"unicode/utf8"
)

type ParserFunc func([]byte) (int, error)

//hack for mutually recursive parsers (see json example)
type box struct {
	f ParserFunc
}

func NewBox() *box {
	return &box{}
}

func (b *box) Wrap(pf ParserFunc) {
	b.f = pf
}

func (b *box) Unwrap() ParserFunc {
	return func(input []byte) (int, error) {
		return b.f(input)
	}
}

func Func(f func(rune) bool) ParserFunc {
	return func(input []byte) (int, error) {
		if len(input) == 0 {
			return 0, EmptyInput
		}

		size := 1
		r := rune(input[0])
		if r >= utf8.RuneSelf {
			r, size = utf8.DecodeRune(input)
		}
		if r == utf8.RuneError {
			return 0, InvalidEncoding
		}
		if !f(r) {
			return 0, WrongRune
		}
		return size, nil
	}
}

func Rune(arg rune) ParserFunc {
	return func(input []byte) (int, error) {
		if len(input) == 0 {
			return 0, EmptyInput
		}

		size := 1
		r := rune(input[0])
		if r >= utf8.RuneSelf {
			r, size = utf8.DecodeRune(input)
		}
		if r == utf8.RuneError {
			return 0, InvalidEncoding
		}
		if r != arg {
			return 0, WrongRune
		}
		return size, nil
	}
}

func Str(s string) ParserFunc {
	return func(input []byte) (int, error) {

		if len(input) == 0 {
			return 0, EmptyInput
		}

		if len(input) >= len(s) && string(input[:len(s)]) == s {
			return len(s), nil
		}

		return 0, NoMatch
	}
}

func EOF(input []byte) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}
	return 0, NonEmptyInput
}

//monadplus reference
func Plus(pf1, pf2 ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		if n, err := pf1(input); err == nil || n > 0 {
			return n, err
		}
		if n, err := pf2(input); err == nil || n > 0 {
			return n, err
		}
		return 0, NoMatch
	}
}

//same as Plus combinator, but for many parser funcs
func Choice(pfs ...ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		for i := range pfs {
			n, err := pfs[i](input)
			if err == nil || n > 0 {
				return n, err
			}
		}
		return 0, NoMatch
	}
}

func Bind(pf1, pf2 ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		var off int

		n, err := pf1(input)
		off += n

		if err != nil {
			return off, err
		}

		n, err = pf2(input[off:])
		off += n

		return off, err
	}
}

func Sequence(pfs ...ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		var off int
		for i := range pfs {
			n, err := pfs[i](input[off:])
			off += n
			if err != nil {
				return off, err
			}
		}
		return off, nil
	}
}

func Try(pf ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		n, err := pf(input)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
}

func CountRange(min, max int, pf ParserFunc) ParserFunc {
	return func(input []byte) (int, error) {
		var off int
		for i := 0; i < max; i++ {
			n, err := pf(input[off:])
			off += n
			if err != nil {
				if i < min {
					return off, err
				}
				return off, nil
			}
		}
		return off, nil
	}
}

func Count(n int, pf ParserFunc) ParserFunc {
	return CountRange(n, n, pf)
}

func Some(pf ParserFunc) ParserFunc {
	return CountRange(1, math.MaxInt64, pf)
}

func Many(pf ParserFunc) ParserFunc {
	return CountRange(0, math.MaxInt64, pf)
}

func Optional(pf ParserFunc) ParserFunc {
	return CountRange(0, 1, pf)
}

func SepBy(val, sep ParserFunc) ParserFunc {
	return Optional(SepBy1(val, sep))
}

func SepBy1(val, sep ParserFunc) ParserFunc {
	return Sequence(val, Many(Bind(sep, val)))
}
