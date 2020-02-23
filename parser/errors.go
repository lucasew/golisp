package parser

import (
	"errors"
)

var (
	ErrEOFWhile                = errors.New("eof when parsing")
	ErrInvalidEntryPoint       = errors.New("invalid entry point")
	ErrPrematureEOF            = errors.New("premature EOF")
	ErrInvalidMarkerRepetition = errors.New("cant use char more than once")
	ErrCantParseAsNumber       = errors.New("cant parse as number")
	ErrInvalidChar             = errors.New("invalid char")
)
