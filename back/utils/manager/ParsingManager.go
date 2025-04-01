package manager

import (
	"fmt"
	"plumlabs/back/utils/lexer"
	"plumlabs/back/utils/parser"
	"plumlabs/back/utils/renderer"
)

/*
//		Parsing manager, take a md file send to renderer, take renderee output send to saveing system of HTML file
//
//		Next step: Add db integration
*/


func ArticleManage(content string) (string, error) {
	
	if content == "" {
		return "", fmt.Errorf("content is empty") 
	}
	lex := lexer.NewLexer(content)
	pars :=	parser.NewParser(lex)
	
	parsevalue := pars.Parse(lexer.EOF)
	if parsevalue == nil{
		return "" , fmt.Errorf("parsing Failed")
	}

	render := renderer.NewRender(parsevalue )
	returnvalue := render.Render(render.Root)	
	if returnvalue == "" {
		return "", fmt.Errorf("render Failed")
	}

	return returnvalue,nil
}

