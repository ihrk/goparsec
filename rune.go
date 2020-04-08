package goparsec

import (
	"unicode/utf8"
)

func RuneFunc(f func(rune) bool) *runeFunc {
	return &runeFunc{f}
}

type runeFunc struct {
	f func(rune) bool
}

func (c *runeFunc) Parse(input []byte) (int, error) {
	if len(input) == 0 {
		return 0, EmptyInput
	}
	size := 1
	r := rune(input[0])
	if r >= utf8.RuneSelf {
		r, size = utf8.DecodeRune(input)
	}
	if !c.f(r) {
		return 0, WrongRune
	}
	return size, nil
}

func (c *runeFunc) Next(input []byte) int {
	return indexFunc(input, c.f, true)
}

func Rune(r rune) *runeFunc {
	return RuneFunc(func(arg rune) bool { return r == arg })
}

func Str(s string) *and {
	runes := []rune(s)
	frs := make([]Combinator, len(runes))
	for i := range runes {
		frs[i] = Rune(runes[i])
	}
	return And(frs...)
}
