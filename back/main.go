
package main

import (
	"fmt"
	utils "plumlabs/back/utils" // Заміни на правильний шлях до лексера
)

func main() {
	input := "![logo](public/images/logo.png)) , [google](https://google.com)"
	lexer := utils.NewLexer(input)

	for tok := lexer.NextToken(); tok.Type != utils.EOF; tok = lexer.NextToken() {
		fmt.Printf("Type: %s, Literal: %s\n", tok.Type, tok.Literal)
	}
}

