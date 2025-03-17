package parser

import (
	"plumlabs/back/utils/lexer"
)

type Parser struct {
	lexer  *lexer.Lexer
	curTok lexer.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	p.curTok = p.lexer.NextToken()
}

func (p *Parser) Parse() *Node {
	root := &Node{Type: "root"}

	for p.curTok.Type != lexer.EOF {
		switch p.curTok.Type {
		case lexer.TEXT:
			root.Children = append(root.Children, p.textParser())
		case lexer.HEADER:
			root.Children = append(root.Children, p.headerParser())
		case lexer.LIST_ITEM:
			root.Children = append(root.Children, p.listParser())
		case lexer.BLOCK_QUOTE:
			root.Children = append(root.Children, p.quoteParser())
		case lexer.CODE_BLOCK:
			root.Children = append(root.Children, p.codeParser())
		case lexer.BOLD:
			root.Children = append(root.Children, p.boldParser())
		case lexer.ITALIC:
			root.Children = append(root.Children, p.italicParser())
		case lexer.STRIKETHROUGH:
			root.Children = append(root.Children, p.strikethroughParser())
		case lexer.AUTO_LINK:
			root.Children = append(root.Children, p.autoLinkParser())
		case lexer.IMAGE:
			root.Children = append(root.Children, p.imageParser())
		default:
			p.NextToken()
		}
	}

	return root
}

func (p *Parser) textParser() *Node {
	node := &Node{Type: TEXT, Value: p.curTok.Literal}
	p.NextToken()
	for p.curTok.Type == lexer.TEXT {
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	return node
}

func (p *Parser) headerParser() *Node {
	node := &Node{Type: HEADER}
	p.NextToken()
	for p.curTok.Type != lexer.NEXT_LINE && p.curTok.Type != lexer.EOF {
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	p.NextToken()
	return node
}

func (p *Parser) listParser() *Node {
	list := &Node{Type: "LIST"}
	for p.curTok.Type == lexer.LIST_ITEM {
		list.Children = append(list.Children, &Node{Type: LIST_ITEM, Value: p.curTok.Literal})
		p.NextToken()
	}
	return list
}

func (p *Parser) quoteParser() *Node {
	node := &Node{Type: BLOCK_QUOTE}
	for p.curTok.Type == lexer.BLOCK_QUOTE {
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	return node
}

func (p *Parser) codeParser() *Node {
	node := &Node{Type: CODE_BLOCK, Value: p.curTok.Literal}
	p.NextToken()
	for p.curTok.Type != lexer.CODE_BLOCK && p.curTok.Type != lexer.EOF {
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	if p.curTok.Type == lexer.CODE_BLOCK {
		p.NextToken()
	}
	return node
}

func (p *Parser) boldParser() *Node {
	node := &Node{Type: BOLD}
	p.NextToken()
	for p.curTok.Type != lexer.BOLD && p.curTok.Type != lexer.EOF {
		node.Children = append(node.Children, &Node{Type: TEXT, Value: p.curTok.Literal})
		p.NextToken()
	}
	p.NextToken()
	return node
}

func (p *Parser) italicParser() *Node {
	node := &Node{Type: ITALIC}
	p.NextToken()
	for p.curTok.Type != lexer.ITALIC && p.curTok.Type != lexer.EOF {
		node.Children = append(node.Children, &Node{Type: TEXT, Value: p.curTok.Literal})
		p.NextToken()
	}
	p.NextToken()
	return node
}

func (p *Parser) strikethroughParser() *Node {
	node := &Node{Type: STRIKETHROUGH}
	p.NextToken()
	for p.curTok.Type != lexer.STRIKETHROUGH && p.curTok.Type != lexer.EOF {
		node.Children = append(node.Children, &Node{Type: TEXT, Value: p.curTok.Literal})
		p.NextToken()
	}
	p.NextToken()
	return node
}

func (p *Parser) autoLinkParser() *Node {
	node := &Node{Type: AUTO_LINK, Value: p.curTok.Literal}
	p.NextToken()
	return node
}

func (p *Parser) imageParser() *Node {
	node := &Node{Type: IMAGE, Value: p.curTok.Literal}
	p.NextToken()
	return node
}
