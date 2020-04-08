package goparsec

//And combinator means sequence of combinators.
//To succeed all combinators must not return any errors.
func And(combs ...Combinator) *and {
	return &and{combs}
}

type and struct {
	combs []Combinator
}

func (c *and) Parse(input []byte) (int, error) {
	off := 0
	for i := range c.combs {
		n, err := c.combs[i].Parse(input[off:])
		off += n
		if err != nil {
			return off, err
		}
	}
	return off, nil
}

func (c *and) Next(input []byte) int {
	if len(c.combs) == 0 {
		return 0
	}

	return c.combs[0].Next(input)
}

//Or combinator means choice of combinators.
//To succeed atleast one of combinators must not return error.
func Or(combs ...Combinator) *or {
	return &or{combs}
}

type or struct {
	combs []Combinator
}

func (c *or) Parse(input []byte) (int, error) {
	for i := range c.combs {
		n, err := c.combs[i].Parse(input)
		if err == nil || n > 0 {
			return n, err
		}
	}
	return 0, NoMatch
}

func (c *or) Next(input []byte) int {
	panic("Next method not implemented for OR combinator")
}

//EOF expects end of input.
func EOF() *eof {
	return &eof{}
}

type eof struct{}

func (c *eof) Parse(input []byte) (int, error) {
	if len(input) != 0 {
		return 0, NonEmptyInput
	}
	return 0, nil
}

func (c *eof) Next(input []byte) int {
	return len(input)
}

func Try(comb Combinator) *try {
	return &try{comb}
}

type try struct {
	comb Combinator
}

func (c *try) Parse(input []byte) (int, error) {
	n, err := c.comb.Parse(input)
	if err != nil {
		return 0, err
	}
	return n, err
}

func (c *try) Next(input []byte) int {
	return c.comb.Next(input)
}
