
package lexer

type TokenType string 

type Token struct{
	Type TokenType
	Literal string
}



const (
	// Blocks
	HEADER        TokenType = "HEADER" // +
	TEXT          TokenType = "TEXT" // + 
	LIST_ITEM     TokenType = "LIST_ITEM" // + 
	BLOCK_QUOTE   TokenType = "BLOCK_QUOTE" // +
	CODE_BLOCK    TokenType = "CODE_BLOCK" // +

	// Text Styling
	BOLD          TokenType = "BOLD" // +
	ITALIC        TokenType = "ITALIC" // +
	STRIKETHROUGH TokenType = "STRIKETHROUGH"// +

	// Interactive elements
	AUTO_LINK     TokenType = "AUTO_LINK"
	IMAGE         TokenType = "IMAGE"

	// Special Tokens
	NEXT_LINE	  TokenType = "NEXT_LINE" // +
	SPACE         TokenType = "SPACE" // +
	TAB           TokenType = "TAB" // +
	EOF           TokenType = "EOF" // + 
	ILLEGAL       TokenType = "ILLEGAL" // +
)


