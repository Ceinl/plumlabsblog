package lexer

type TokenType string 

type Token struct{
	Type TokenType
	Literal string
}



const (

	HEADER        TokenType = "HEADER" //Token "#" or HTML block <h1>  
	TEXT          TokenType = "TEXT" // Token "..." or HTML block <p> 
	LIST_ITEM     TokenType = "LIST_ITEM" // Token "*" or HTML block <li> 
	BLOCK_QUOTE   TokenType = "BLOCK_QUOTE" // Token ">" or HTML block <blockquote>
	CODE_BLOCK    TokenType = "CODE_BLOCK" // Token "```" or HTML <pre><code>...</code></pre>"

	BOLD          TokenType = "BOLD" // ** or HTML block <strong> 
	ITALIC        TokenType = "ITALIC" // * or HTML block <em>
	STRIKETHROUGH TokenType = "STRIKETHROUGH"// ~~ or HTML block <del>

	AUTO_LINK     TokenType = "AUTO_LINK" // Token "[text](url)" or HTML block <a>
	IMAGE         TokenType = "IMAGE" // Token "![text](url)" or HTML block <img>

	NEXT_LINE	  TokenType = "NEXT_LINE" // New Line character
	SPACE         TokenType = "SPACE" // Whitespace character
	TAB           TokenType = "TAB" // Tab character
	EOF           TokenType = "EOF" // End Of File 
	ILLEGAL       TokenType = "ILLEGAL" // Not recognised token
)


