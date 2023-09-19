package main

import (
	"fmt"

	"github.com/AndreiAlbert/brainfuckgo/generators"
)

func main() {
	i := generators.New(",[.,]")
	result := i.Evaluate()
	fmt.Println(result.String())
}
