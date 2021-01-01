package main

import (
	"fmt"
	"github.com/selectiveduplicate/cosmoc/lexer"
)

func main() {
	for {
		var input string
		fmt.Print("Welcome to Cosmoc v0.1a, the lamest thing ever!\ncalc>>")
		fmt.Scanln(&input)
		interpreter := lexer.NewInterpreter(input, 0, nil, string(input[0]))
		output := interpreter.Expression()
		fmt.Println(output)
	}
}
