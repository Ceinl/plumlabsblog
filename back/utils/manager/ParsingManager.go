package manager

import (
	"fmt"
	"plumlabs/back/utils/lexer"
	"plumlabs/back/utils/parser"
	"plumlabs/back/utils/renderer"
)

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

