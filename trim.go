package goparsec

import (
	"unicode"
	"unicode/utf8"
)

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// from standart lib "bytes"
// but instead of -1 it returns len of slice
func indexFunc(s []byte, f func(r rune) bool, truth bool) int {
	start := 0
	for start < len(s) {
		wid := 1
		r := rune(s[start])
		if r >= utf8.RuneSelf {
			r, wid = utf8.DecodeRune(s[start:])
		}
		if f(r) == truth {
			return start
		}
		start += wid
	}
	return len(s)
}

//TrimSpace is same as TrimFunc(unicode.IsSpace)
func TrimSpace(input []byte) (int, error) {
	start := 0
	for ; start < len(input); start++ {
		c := input[start]
		if c >= utf8.RuneSelf {
			start += indexFunc(input[start:], unicode.IsSpace, false)
			return start, nil
		}
		if asciiSpace[c] == 0 {
			break
		}
	}
	return start, nil
}

//TrimFunc is same as Many(Func)
func TrimFunc(f func(r rune) bool) ParserFunc {
	return func(input []byte) (int, error) {
		return indexFunc(input, f, false), nil
	}
}

//TrimFunc1 is same as Some(Func)
func TrimFunc1(f func(r rune) bool) ParserFunc {
	return func(input []byte) (int, error) {
		n := indexFunc(input, f, false)
		if n == 0 {
			return 0, EmptyInput
		}
		return n, nil
	}
}
