package renderer

import (
	"plumlabs/back/utils/parser"
	"strings"
)
type Renderer struct{
    Root *parser.Node
}
func NewRender(root *parser.Node) *Renderer {
    return &Renderer{Root: root}
}
func (r *Renderer)Render(node *parser.Node) string {
    if node == nil {
        return ""
    }
    var html strings.Builder 

    if node.Value != "" && len(node.Children) == 0 {
        return node.Value
    }

    switch node.Type{
    case parser.DOCUMENT: 
        for _, child := range node.Children{
            html.WriteString(r.Render(child))
        }

    case parser.HEADER:
        html.WriteString("<h1>")
        for _, child := range node.Children{
            html.WriteString(r.Render(child))
        }
        html.WriteString("</h1>")

    case parser.TEXT:
        if len(node.Children) > 0 {
            for _, child := range node.Children{
                html.WriteString(r.Render(child))
            }
        } else {
            html.WriteString(node.Value)
        }

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
        if len(node.Children) > 0 {
            for _, child := range node.Children {
                html.WriteString(r.Render(child))
            }
        } else if node.Value != "" {
            html.WriteString(node.Value)
        } else {
            html.WriteString("<p>Error</p>")
        }
    } 
    return html.String() 
}
func (r *Renderer)imageRender(value string) string {
    parts := strings.Split(value ," -> ")
    if len(parts) != 2 {
        //return "<img src=\"\" alt=\"Invalid image format\">"
		return "<p> Images currently not supported:</p>"
    }

	return "<p> Images currently not supported:</p>"
    //return "<img src=\"" + parts[1] + "\" alt=\"" + parts[0] + "\">"
}
func (r *Renderer)linkRenderer(value string) string{
    parts := strings.Split(value ," -> ")
    if len(parts) != 2 {
        return "<a href=\"\">Invalid link format</a>"
    }
    return "<a href=\"" + parts[1] + "\">" + parts[0] + "</a>"
}
