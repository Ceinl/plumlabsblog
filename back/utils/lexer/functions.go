package lexer


func (l *Lexer) newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal:string(ch)}
}

// Lexer funcs
func (l *Lexer) readText() string {
	possition := l.position 
	for isLetter(l.ch) || isDigit(l.ch) || isSymbol(l.ch) {
		l.readChar()
	}
	return l.input[possition:l.position]
	
}

func (l *Lexer) readChar(){
	if l.readPosition >= len(l.input){
		l.ch = 0
	}else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte{
	if l.readPosition >= len(l.input){
		return 0
	}else{
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekNextChar() byte{
	if l.readPosition+1 >= len(l.input){
		return 0
	}else{
		return l.input[l.readPosition+1]
	}
}


func (l *Lexer) readLinkURL() string {
	if l.ch != '(' {
		return ""
	}
	l.readChar() 

	start := l.position
	for l.ch != ')' {
		l.readChar()
	}

	url := l.input[start:l.position]
	l.readChar()

	return url
}


func (l *Lexer) readLinkText() string {
	l.readChar() 
	start := l.position
	for l.ch != ']' && l.ch != 0 {
		l.readChar()
	}
	text := l.input[start:l.position]
	l.readChar()
	return text
}

//-------

// Other funcs
func isLetter(ch byte)bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch ==',' || ch =='.' || ch ==' '
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSymbol(ch byte) bool {
	return ch == '!' || ch == '?' || ch == ',' || ch == '.' || ch == ':' || ch == ';' 
}

// -------

