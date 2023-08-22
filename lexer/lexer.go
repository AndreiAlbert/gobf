package lexer

import (
	"strings"
)

const (
	EOF = "EOF"

	INC_POINTER = ">"
	DEC_POINTER = "<"

	INC_VALUE = "+"
	DEC_VALUE = "-"

	LOOP_START = "["
	LOOP_END   = "]"

	INPUT  = ","
	OUTPUT = "."
)

type tokenType string

type Token struct {
	Type         tokenType
	LiteralValue string
}

type Lexer struct {
	Input        string
	currPosition int
	known        map[string]tokenType
}

func New(input string) *Lexer {
	lex := &Lexer{Input: input, known: make(map[string]tokenType)}
	lex.Input = strings.ReplaceAll(lex.Input, "\n", "")
	lex.Input = strings.ReplaceAll(lex.Input, "\t", "")
	lex.known["+"] = INC_VALUE
	lex.known["-"] = DEC_VALUE
	lex.known["<"] = DEC_POINTER
	lex.known[">"] = INC_POINTER
	lex.known["["] = LOOP_START
	lex.known["]"] = LOOP_END
	lex.known["."] = OUTPUT
	lex.known[","] = INPUT
	return lex
}

func (l *Lexer) GetTokens() []Token {
	var tokens []Token
	for l.currPosition < len(l.Input) {
		charAtPositiion := string(l.Input[l.currPosition])
		tokenType, exists := l.known[charAtPositiion]
		if !exists {
			l.currPosition++
			continue
		}
		tokens = append(tokens, Token{Type: tokenType, LiteralValue: charAtPositiion})
		l.currPosition++
	}
	return tokens
}
