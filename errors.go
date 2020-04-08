package goparsec

import (
	"errors"
)

var EmptyInput = errors.New("unexpected end of input")
var NonEmptyInput = errors.New("expected end of input")
var InvalidEncoding = errors.New("invalid encoding")
var WrongRune = errors.New("got incorrect symbol")
var NoMatch = errors.New("no combinators matched")
