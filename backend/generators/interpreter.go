package generators

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/AndreiAlbert/brainf/lexer"
)

type Interpreter struct {
	tokens        []lexer.Token
	memory        [3000]uint8
	memoryPointer int
	tokenPointer  int
}

func New(input string) *Interpreter {
	lex := lexer.New(input)
	tokens := lex.GetTokens()
	return &Interpreter{tokens: tokens}
}

func (i *Interpreter) Evaluate() (strings.Builder, error) {
	var str strings.Builder
	loopStack := []int{}
	for i.tokenPointer < len(i.tokens) {
		token := i.tokens[i.tokenPointer]
		switch token.LiteralValue {
		case lexer.INC_VALUE:
			i.memory[i.memoryPointer]++
		case lexer.DEC_VALUE:
			i.memory[i.memoryPointer]--
		case lexer.INC_POINTER:
			i.memoryPointer++
		case lexer.DEC_POINTER:
			i.memoryPointer--
		case lexer.OUTPUT:
			str.WriteByte(i.memory[i.memoryPointer])
		case lexer.INPUT:
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()
			if err != nil {
				log.Fatal("could not read from stdin")
			}
			i.memory[i.memoryPointer] = uint8(char)
		case lexer.LOOP_START:
			if i.memory[i.memoryPointer] == 0 {
				loopDepth := 1
				for idx := i.tokenPointer + 1; idx < len(i.tokens); idx++ {
					loopToken := i.tokens[idx]
					if loopToken.Type == lexer.LOOP_START {
						loopDepth++
					} else if loopToken.Type == lexer.LOOP_END {
						loopDepth--
						if loopDepth == 0 {
							i.tokenPointer = idx
							break
						}
					}
				}
				if loopDepth != 0 {
					return strings.Builder{}, errors.New("No close bracked")
				}
			} else {
				loopStack = append(loopStack, i.tokenPointer)
			}
		case lexer.LOOP_END:
			if i.memory[i.memoryPointer] != 0 {
				if len(loopStack) == 0 {
					return strings.Builder{}, errors.New("No close bracked")
				}
				i.tokenPointer = loopStack[len(loopStack)-1]
			} else {
				if len(loopStack) == 0 {
					return strings.Builder{}, errors.New("No close bracked")
				}
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
		i.tokenPointer++
	}
	return str, nil
}
