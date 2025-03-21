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
			listblock := p.parseListBlock()
			root.Children = append(root.Children, listblock)
			p.NextToken()
		case lexer.BLOCK_QUOTE:
			blockNode := &Node{Type: NodeType(lexer.BLOCK_QUOTE)}
			p.NextToken()
			blockNode.Children = p.Parse(lexer.NEXT_LINE).Children
			root.Children = append(root.Children, blockNode)

		case lexer.CODE_BLOCK: 
			codeNode := p.CodeParser() 
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

		case lexer.IMAGE:
			imageNode := &Node{Type: NodeType(lexer.IMAGE), Value: p.curTok.Literal}
			p.NextToken()
			imageNode.Children = p.Parse(lexer.NEXT_LINE).Children
			root.Children = append(root.Children, imageNode)

		default:
			p.NextToken()
		}
	}
	return root
}

func (p *Parser) nlParser() *Node {
	node := &Node{Type: NEXT_LINE}
	p.NextToken()
	return node
}

func (p *Parser) textParser() *Node {
	node := &Node{Type: TEXT, Value: p.curTok.Literal}
	p.NextToken()
	for p.curTok.Type == lexer.TEXT || p.curTok.Type == lexer.SPACE || p.curTok.Type == lexer.TAB {
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	return node
}

func (p *Parser) CodeParser() *Node {
	node := &Node{Type: CODE_BLOCK, Value: ""}
	p.NextToken()
	for p.curTok.Type != lexer.CODE_BLOCK && p.curTok.Type != lexer.EOF{
		node.Value += p.curTok.Literal
		p.NextToken()
	}
	if p.curTok.Type == lexer.CODE_BLOCK{
		p.NextToken()
	}
	return node
}

func (p *Parser)parseListBlock() *Node{

	listblock := &Node{Type: LIST_BLOCK}
	
	for p.curTok.Type == lexer.LIST_ITEM{

		p.NextToken()
		listitem := &Node{Type: LIST_ITEM}
		listitem.Children = p.Parse(lexer.NEXT_LINE).Children
		listblock.Children = append(listblock.Children, listitem)

		if p.curTok.Type != lexer.LIST_ITEM && p.curTok.Type != lexer.NEXT_LINE{
			break
		}
	}

	return listblock
}



















