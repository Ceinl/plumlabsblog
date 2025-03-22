package main

import (
	"fmt"
	"plumlabs/back/utils/lexer"
	"plumlabs/back/utils/parser"
	"plumlabs/back/utils/renderer" // переконайтеся, що шлях імпорту правильний
)
func main() {
	// Тест для лексера і парсера (ваш існуючий код)
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
	
	// Тест для рендереру
	testRenderer()
}

// Функція для тестування рендереру
func testRenderer() {
	fmt.Println("\n\n===== Renderer Tests =====")
	
	// Тест 1: Простий текст
	testRenderCase("Simple text", "This is simple text.")
	
	// Тест 2: Заголовок
	testRenderCase("Header", "# This is a header")
	
	// Тест 3: Форматований текст
	testRenderCase("Formatted text", "This is **bold** and *italic* text.")
	
	// Тест 4: Списки
	testRenderCase("Lists", "- Item 1\n- Item 2\n- Item 3")
	
	// Тест 5: Вкладені списки
	testRenderCase("Nested lists", "- Item 1\n  - Nested item 1\n  - Nested item 2\n- Item 2")
	
	// Тест 6: Блок коду
	testRenderCase("Code block", "```\nfunc example() {\n  return true\n}\n```")
	
	// Тест 7: Цитата
//	testRenderCase("Blockquote", "> This is a quoted text.")
	
	// Тест 8: Зображення
//	testRenderCase("Image", "![Alt text] -> http://example.com/image.jpg")
	
	// Тест 9: Посилання
//	testRenderCase("Link", "[Link text] -> http://example.com")
	
	// Тест 10: Комбінований документ
//	testRenderCase("Combined document", "# Document Title\n\nThis is a paragraph with **bold** and *italic* text.\n\n- List item 1\n- List item 2\n\n> A quote\n\n![Image] -> http://example.com/img.jpg")
}

// Допоміжна функція для тестування окремих випадків рендерингу
func testRenderCase(testName string, input string) {
	fmt.Printf("\n--- Test: %s ---\n", testName)
	fmt.Println("Input:")
	fmt.Println(input)
	
	// Lexing та parsing
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	root := p.Parse(lexer.EOF)
	
	// Rendering
	r := renderer.NewRender(root)
	html := r.Render(root)
	
	fmt.Println("\nRendered HTML:")
	fmt.Println(html)
	fmt.Println("---------------")
}

// Ваші існуючі функції printTokens та printNode
func printTokens(l *lexer.Lexer) {
	for {
		tok := l.NextToken()
		fmt.Printf("Token Type: %s, Literal: %q\n", tok.Type, tok.Literal)
		if tok.Type == lexer.EOF {
			break
		}
	}
}

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
