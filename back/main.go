package main

import (
	"fmt"
	"plumlabs/back/utils/lexer"
	"plumlabs/back/utils/parser"
)

func main() {
	input := "-asd\n - ** asd ** \n -asd\n"
	
	// Print the original input
	fmt.Println("Original Input:")
	fmt.Printf("%q\n\n", input)
	
	// Print lexer tokens first
	fmt.Println("Lexer Results:")
	l := lexer.NewLexer(input)
	printTokens(l)
	
	// Reset lexer for parser
	l = lexer.NewLexer(input)
	
	// Create parser and run parsing
	p := parser.NewParser(l)
	root := p.Parse(lexer.EOF)
	
	// Print parsing result
	fmt.Println("\nParser Results:")
	printNode(root, 0)
}

// Function to print all tokens from the lexer
func printTokens(l *lexer.Lexer) {
	for {
		tok := l.NextToken()
		fmt.Printf("Token Type: %s, Literal: %q\n", tok.Type, tok.Literal)
		if tok.Type == lexer.EOF {
			break
		}
	}
}

// Function for recursive node tree printing
func printNode(node *parser.Node, indent int) {
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += "  "
	}
	fmt.Printf("%sNode Type: %s, Value: %s\n", indentation, node.Type, node.Value)
	for _, child := range node.Children {
		printNode(child, indent+1)
	}
}
