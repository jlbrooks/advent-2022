package main

import (
	"fmt"
	"strconv"

	"github.com/jlbrooks/advent-2022/shared"
)

type tokenType uint8

const (
	prompt tokenType = iota
	commandLs
	commandDir
	stringLiteral
	numericLiteral
)

type token struct {
	typ      tokenType
	value    string
	intValue int
}

func newNumericToken(intValue int) token {
	return token{
		typ:      numericLiteral,
		intValue: intValue,
	}
}

func newStringToken(value string) token {
	return token{
		typ:   stringLiteral,
		value: value,
	}
}

type lexerState uint8

const (
	start lexerState = iota
	parseCommand
)

type commandLexer struct {
	state  lexerState
	reader shared.Reader
}

func newCommandLexer(r shared.Reader) *commandLexer {
	return &commandLexer{
		state:  start,
		reader: r,
	}
}

func (l *commandLexer) next() token {
	switch l.state {
	case start:
		return l.processStart()
	}

	panic("uh oh!")
}

func (l *commandLexer) processStart() token {
	next, err := l.reader.PeekNext()
	if err != nil {
		panic(err)
	}

	switch next {
	case '$':
		l.reader.ConsumeNext()
		l.state = parseCommand
		return token{
			typ: prompt,
		}
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':
		return l.processNumericLiteral()
	default:
		return l.processStringLiteral()
	}

	panic("uh oh!")
}

func (l *commandLexer) processNumericLiteral() token {
	var digits []rune
	for next, err := l.reader.ConsumeNext(); err != nil && next != ' '; {
		digits = append(digits, next)
	}
	num, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(fmt.Sprintf("Error parsing numeric literal %s", string(digits)))
	}

	return newNumericToken(num)
}

func (l *commandLexer) processStringLiteral() token {
	var chars []rune
	for next, err := l.reader.ConsumeNext(); err != nil && next != ' '; {
		chars = append(chars, next)
	}

	return newStringToken(string(chars))
}

func main() {
	input := shared.ReadContinuous("day07/example.txt")
	reader := shared.NewReader(input)
	lexer := newCommandLexer(reader)
}
