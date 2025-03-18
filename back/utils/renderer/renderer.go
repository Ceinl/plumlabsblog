package renderer

import (
	"fmt"
	"plumlabs/back/utils/parser"
	"strings"
)

/*

const (
	// Blocks
	HEADER        NodeType = "HEADER" // +
	TEXT          NodeType = "TEXT" // +
	LIST_ITEM     NodeType = "LIST_ITEM" // +
	BLOCK_QUOTE   NodeType = "BLOCK_QUOTE" // +
	CODE_BLOCK    NodeType = "CODE_BLOCK" // +

	// Text Styling
	BOLD          NodeType = "BOLD" //+
	ITALIC        NodeType = "ITALIC" // +
	STRIKETHROUGH NodeType = "STRIKETHROUGH" // +

	// Interactive elements
	AUTO_LINK     NodeType = "AUTO_LINK" // +
	IMAGE         NodeType = "IMAGE" // +

	// Special Tokens
	NEXT_LINE     NodeType = "NEXT_LINE" // +
	SPACE         NodeType = "SPACE"
	TAB           NodeType = "TAB" // +
	EOF           NodeType = "EOF"
	ILLEGAL       NodeType = "ILLEGAL"
)


*/

func renderer(node *parser.Node) string{

	switch node.Type {
	case parser.HEADER:
		return fmt.Sprintf("<h1> %s </h1",node.Value)
	case parser.TEXT:
		return fmt.Sprintf("<p> %s </p>",node.Value)
	case parser.LIST_ITEM:
		// Render list item
	case parser.BLOCK_QUOTE:
		// Render block quote
		return fmt.Sprintf("<blockquote> %s </blockquote>", node.Value)
	case parser.CODE_BLOCK:
		// Render code block
	case parser.BOLD:
		// Render bold text
		return fmt.Sprintf("<b> %s </b>", node.Value)
	case parser.ITALIC:
		// Render italic text
	case parser.STRIKETHROUGH:
		// Render strikethrough text
	case parser.AUTO_LINK:
		// Render auto link
	case parser.IMAGE:
		// Render image
	case parser.NEXT_LINE:
		// Render next line
	case parser.SPACE:
		// Render space
	case parser.TAB:
		// Render tab
	case parser.EOF:
		// Render EOF
	case parser.ILLEGAL:
		// Render illegal
	default:
	return ""
	}
	return ""
}

func render_children(node *parser.Node){
	var sb strings.Builder
	for _, child := range node.Children {
	sb.WriteString(renderer(child))
	}
}
