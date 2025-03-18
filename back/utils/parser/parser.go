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

func (p *Parser) Parse(endToken lexer.TokenType) *Node {

	root := &Node{Type: NodeType(endToken)}

	for p.curTok.Type != endToken && p.curTok.Type != lexer.EOF {
		switch p.curTok.Type {
			case lexer.TEXT:
				root.Children = append(root.Children, p.textParser())
			case lexer.HEADER:
			    headerNode := &Node{Type: NodeType(lexer.HEADER)}
			    p.NextToken() 
			    headerNode.Children = p.Parse(lexer.NEXT_LINE).Children
			    root.Children = append(root.Children, headerNode)
			case lexer.LIST_ITEM:
				listNode := &Node{Type: NodeType(lexer.LIST_ITEM)}
				p.NextToken()
				listNode.Children = p.Parse(lexer.NEXT_LINE).Children
				root.Children = append(root.Children, listNode)
			case lexer.BLOCK_QUOTE:
				blockNode := &Node{Type: NodeType(lexer.BLOCK_QUOTE)}
				p.NextToken()
				blockNode.Children = p.Parse(lexer.NEXT_LINE).Children
				root.Children = append(root.Children, blockNode)
			case lexer.CODE_BLOCK:
				codeNode := &Node{Type: NodeType(lexer.CODE_BLOCK)}
				p.NextToken()
				codeNode.Children = p.Parse(lexer.CODE_BLOCK).Children
				root.Children = append(root.Children, codeNode)
			case lexer.BOLD:
				codeNode := &Node{Type: NodeType(lexer.BOLD)}
				p.NextToken()
				codeNode.Children = p.Parse(lexer.BOLD).Children
				root.Children = append(root.Children, codeNode)
			case lexer.ITALIC:
				codeNode := &Node{Type: NodeType(lexer.ITALIC)}
				p.NextToken()
				codeNode.Children = p.Parse(lexer.ITALIC).Children
				root.Children = append(root.Children, codeNode)
			case lexer.STRIKETHROUGH:
				codeNode := &Node{Type: NodeType(lexer.STRIKETHROUGH)}
				p.NextToken()
				codeNode.Children = p.Parse(lexer.STRIKETHROUGH).Children
				root.Children = append(root.Children, codeNode)
			case lexer.AUTO_LINK:
				linkNode := &Node{Type: NodeType(lexer.AUTO_LINK), Value: p.curTok.Literal}
				p.NextToken()
				linkNode.Children = p.Parse(lexer.NEXT_LINE).Children
				root.Children = append(root.Children, linkNode)
			case lexer.NEXT_LINE:
				root.Children = append(root.Children, p.nlParser())
			case lexer.TAB:
				root.Children = append(root.Children, p.tabParser())
			case lexer.IMAGE:
				imageNode := &Node{Type: NodeType(lexer.IMAGE), Value: p.curTok.Literal}
				p.NextToken()
				imageNode.Children = p.Parse(lexer.NEXT_LINE).Children
				root.Children = append(root.Children, imageNode)
			case lexer.SPACE:
				root.Children = append(root.Children, p.spaceParser())
			default:
				p.NextToken()
		}
	}
	return root

}

func (p *Parser) Parse3() *Node {
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
		case lexer.NEXT_LINE:
			root.Children = append(root.Children, p.nlParser())
		case lexer.TAB:
			root.Children = append(root.Children, p.tabParser())
		case lexer.IMAGE:
			root.Children = append(root.Children, p.imageParser())
		case lexer.SPACE:
			root.Children = append(root.Children, p.spaceParser())
		default:
			p.NextToken()
		}
	}

	return root
}

func (p *Parser) spaceParser() *Node {
	node := &Node{Type: SPACE}
	p.NextToken()
	return node
}


func (p *Parser) tabParser() *Node {
	node := &Node{Type: TAB}
	p.NextToken()
	return node
}


func (p *Parser) nlParser() *Node {
	node := &Node{Type: NEXT_LINE}
	p.NextToken()
	return node
}


func (p *Parser) textParser() *Node {
	node := &Node{Type: TEXT, Value: p.curTok.Literal}
	p.NextToken()
	for p.curTok.Type == lexer.TEXT /*|| p.curTok.Type == lexer.SPACE || p.curTok.Type == lexer.TAB */{
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
