package parser

type NodeType string

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

type Node struct{
	Type NodeType
	Value string
	Children []*Node
}
