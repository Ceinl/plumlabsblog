package renderer

/*
TODO: fix and make a list work correctly


*/

import (
	"plumlabs/back/utils/parser"
	"strings"
)

type Renderer struct{
	root *parser.Node
}

func NewRender(root *parser.Node) *Renderer {
	return &Renderer{root: root}
}

func (r *Renderer)Render(node *parser.Node) string {
	if node == nil {
		return ""
	}

	var html strings.Builder 
	
	switch node.Type{
	case parser.HEADER:
		html.WriteString("<h1>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</h1>")

	case parser.TEXT:
		html.WriteString("<p>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</p>")

	case parser.BOLD:
		html.WriteString("<b>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</b>")

	case parser.ITALIC:
		html.WriteString("<i>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</i>")

	case parser.STRIKETHROUGH:
		html.WriteString("<strike>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</strike>")

	case parser.LIST_BLOCK:
		html.WriteString("<ul>")
		 for _, child := range node.Children{
			html.WriteString(r.Render(child))
		 }
		html.WriteString("</ul>")

	case parser.LIST_ITEM:
		html.WriteString("<li>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</li>")

	case parser.CODE_BLOCK:
		html.WriteString("<pre><code>")
		html.WriteString(node.Value)
		html.WriteString("</code></pre>")

	case parser.BLOCK_QUOTE:
		html.WriteString("<blockquote>")
		for _, child := range node.Children{
			html.WriteString(r.Render(child))
		}
		html.WriteString("</blockquote>")

	case parser.IMAGE:
		html.WriteString(r.imageRender(node.Value))

	case parser.AUTO_LINK:
		html.WriteString(r.linkRenderer(node.Value))
	
	default:
		html.WriteString("<p>")
			html.WriteString("Error")
		html.WriteString("</p>")
	} 
	return html.String() 
}



func (r *Renderer)imageRender(value string) string {
	parts := strings.Split(value ," -> ")
	return "<img src=\"" + parts[1] + "\" alt=\"" + parts[0] + "\">"
}

func (r *Renderer)linkRenderer(value string) string{
	parts := strings.Split(value ," -> ")
	return "<a href=\"" + parts[1] + "\">" + parts[0] + "</a>"
}











