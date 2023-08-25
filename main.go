package main

import (
	"github.com/AndreiAlbert/brainfuckgo/generators"
)

func main() {
	i := generators.New(",.")
	i.Evaluate()
}
