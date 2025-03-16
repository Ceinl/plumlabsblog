package lexer

type Lexer struct{
	input string
	position int 
	readPosition int 
	ch byte
}


func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// GET A TOKEN
func (l *Lexer) NextToken() Token {
	var tok Token
	
	switch l.ch{
	// BLOCK 
	case '#': 
	tok = l.newToken(HEADER,  l.ch)

	case '*':
		if l.peekChar() == '*'{ 
			l.readChar()
			l.readChar()
			tok = l.newToken(BOLD, l.ch)
		}else{
			tok = l.newToken(LIST_ITEM, l.ch)
		}

	case '>':
		tok = l.newToken(BLOCK_QUOTE,l.ch)
	
	case '`':
		if l.peekChar() == '`' && l.peekNextChar() == '`'{
			tok = l.newToken(CODE_BLOCK, l.ch)
		}

	// STYLING
	case '_':
		if l.peekChar() == '_' && l.peekNextChar()== '_' {
			l.readChar()
			l.readChar()
			tok = l.newToken(ITALIC,l.ch)
		}

	case '~':
		if l.peekChar() == '~' && l.peekNextChar() == '~'{
			tok = l.newToken(STRIKETHROUGH,l.ch)
		}

	// INTERACTIVE ELEMENTS 
	case '[':
		
		text := l.readLinkText()
		url := l.readLinkURL()
		tok = Token{Type: AUTO_LINK, Literal: text + " -> " + url}

	case '!':
		if l.peekChar() == '['{
			l.readChar()
			text := l.readLinkText()
			url := l.readLinkURL()
			tok = Token{Type: IMAGE, Literal: text+ " -> " + url}
		}
	// SPECIAL TOKENS

	case ' ':
		tok = l.newToken(SPACE,l.ch)
	case '\n':
		tok = l.newToken(NEXT_LINE,l.ch)
	case '\t':
		tok = l.newToken(TAB,l.ch)
	case 0: 
	tok = l.newToken(EOF, l.ch)

	default:
	if isLetter(l.ch){
		text := l.readText()
		return Token{Type: TEXT, Literal: text}
	}
	
	tok = l.newToken(ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

