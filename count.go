package goparsec

import (
	"math"
)

func Count(n int, comb Combinator) *count {
	return Count2(n, n, comb)
}

func Count2(min, max int, comb Combinator) *count {
	return &count{min, max, comb}
}

type count struct {
	min  int
	max  int
	comb Combinator
}

func (c *count) Parse(input []byte) (int, error) {
	off := 0
	for i := 0; i < c.max; i++ {
		n, err := c.comb.Parse(input[off:])
		if err != nil {
			if n > 0 {
				return off + n, err
			}
			if i < c.min {
				return off, err
			}
		}
		if n == 0 {
			return off, nil
		}
		off += n
	}

	return off, nil
}

func (c *count) Next(input []byte) int {
	if c.min == 0 {
		return 0
	}
	return c.comb.Next(input)
}

func Some(comb Combinator) *count {
	return Count2(1, math.MaxInt64, comb)
}

func Many(comb Combinator) *count {
	return Count2(0, math.MaxInt64, comb)
}

func Opt(comb Combinator) *count {
	return Count2(0, 1, comb)
}

func SepBy(val, sep Combinator) *count {
	return Opt(And(val, Many(And(sep, val))))
}
