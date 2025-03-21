package main

import (
	"fmt"
	"plumlabs/back/utils/lexer"
	"plumlabs/back/utils/parser"
)

func main() {

	input := "[url](text asd asd)"	
	// Створення лексера
	l := lexer.NewLexer(input)

	// Створення парсера
	p := parser.NewParser(l)

	// Запуск парсингу
	root := p.Parse(lexer.EOF)

	// Виведення результату
	printNode(root, 0)
}

// Функція для рекурсивного виведення дерева нод
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

