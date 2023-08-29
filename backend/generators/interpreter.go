package generators

import (
	"log"
	"os"
	"strings"

	"github.com/AndreiAlbert/brainfuckgo/lexer"
)

type Interpreter struct {
	tokens      []lexer.Token
	memory      [3000]int
	currPointer int
}

func New(input string) *Interpreter {
	lex := lexer.New(input)
	tokens := lex.GetTokens()
	return &Interpreter{tokens: tokens}
}

func (i *Interpreter) Evaluate() strings.Builder {
	var str strings.Builder
	for _, token := range i.tokens {
		switch token.LiteralValue {
		case lexer.INC_VALUE:
			i.memory[i.currPointer]++
		case lexer.DEC_VALUE:
			i.memory[i.currPointer]--
		case lexer.INC_POINTER:
			i.currPointer++
		case lexer.DEC_POINTER:
			i.currPointer--
		case lexer.OUTPUT:
			str.WriteString(string(rune(i.memory[i.currPointer])))
		case lexer.INPUT:
			buf := make([]byte, 1)
			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal("Could not read standard input")
			}
			i.memory[i.currPointer] = int(buf[0])
		}
	}
	return str
}
