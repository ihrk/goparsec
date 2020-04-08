package goparsec

import (
	"unicode/utf8"
)

//Combinator is main parsing interface
type Combinator interface {
	//Parse returns number of matching bytes
	//and possible error while parsing
	//possible results:
	//1) nonzero, nil - parsed successfully;
	//2) n, error - parsed with error, number of bytes parsed must be ignored;
	//3) 0, nil - failed to parse, but parsing was optional.
	Parse([]byte) (int, error)
	//Next returns number bytes that needs to be skipped to
	//reach likely matching result.
	Next([]byte) int
}

func skipRune(input []byte) int {
	if len(input) == 0 {
		return 0
	}
	size := 1
	if input[0] >= utf8.RuneSelf {
		_, size = utf8.DecodeRune(input)
	}
	return size
}

func FindAll(input []byte, comb Combinator) (res [][]byte) {
	off := 0
	for off < len(input) {
		off += comb.Next(input[off:])
		n, err := comb.Parse(input[off:])
		if err != nil || n == 0 {
			off += skipRune(input[off:])
			continue
		}
		res = append(res, input[off:off+n])
		if n == 0 {
			off += skipRune(input[off:])
			continue
		}
		off += n
	}
	return
}

func FindAllString(input []byte, comb Combinator) (res []string) {
	off := 0
	for off < len(input) {
		off += comb.Next(input[off:])
		n, err := comb.Parse(input[off:])
		if err != nil {
			off += skipRune(input[off:])
			continue
		}
		res = append(res, string(input[off:off+n]))
		if n == 0 {
			off += skipRune(input[off:])
			continue
		}
		off += n
	}
	return
}
