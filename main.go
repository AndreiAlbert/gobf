package main

import (
	"fmt"

	"github.com/AndreiAlbert/brainfuckgo/lexer"
)

func main() {
	l := lexer.New(">>++--[].abcd[[[]]]")
	tokens := l.GetTokens()
	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}
}
