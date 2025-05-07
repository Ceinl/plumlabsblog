package renderer

import (
	"html"
	"plumlabs/back/utils/parser"
	"strings"
)

type Renderer struct {
	Root *parser.Node
}

func NewRender(root *parser.Node) *Renderer {
	return &Renderer{Root: root}
}

func (r *Renderer) Render(node *parser.Node) string {
	if node == nil {
		return ""
	}
	var builder strings.Builder

	switch node.Type {
	case parser.DOCUMENT:
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}

	case parser.HEADER:
		builder.WriteString(`<h1 class="text-4xl font-bold mb-4">`)
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</h1>")

	case parser.TEXT:
		builder.WriteString(html.EscapeString(node.Value))

	case parser.BOLD:
		builder.WriteString("<b>")
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</b>")

	case parser.ITALIC:
		builder.WriteString("<i>")
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</i>")

	case parser.STRIKETHROUGH:
		builder.WriteString("<del>")
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</del>")

	case parser.LIST_BLOCK:
		builder.WriteString(`<ul class="list-disc pl-6 space-y-1">`)
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</ul>")

	case parser.LIST_ITEM:
		builder.WriteString(`<li class="text-base">`)
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</li>")

	case parser.CODE_BLOCK:
		builder.WriteString("<pre><code>")
		builder.WriteString(html.EscapeString(node.Value))
		builder.WriteString("</code></pre>")

	case parser.BLOCK_QUOTE:
		builder.WriteString("<blockquote>")
		for _, child := range node.Children {
			builder.WriteString(r.Render(child))
		}
		builder.WriteString("</blockquote>")

	case parser.IMAGE:
		builder.WriteString(r.imageRender(node.Value))
	case parser.AUTO_LINK:
		builder.WriteString(r.linkRenderer(node.Value))
	case parser.NEXT_LINE:
		builder.WriteString("<br>")

	default:
		if len(node.Children) > 0 {
			for _, child := range node.Children {
				builder.WriteString(r.Render(child))
			}
		} else if node.Value != "" {
			builder.WriteString("<p>")
			builder.WriteString(html.EscapeString(node.Value))
			builder.WriteString("</p>")
		}
	}
	return builder.String()
}

func (r *Renderer) imageRender(value string) string {
	parts := strings.Split(value, " -> ")
	if len(parts) != 2 {
		return "<img src=\"\" alt=\"Invalid image format\">"
	}

	altText := html.EscapeString(parts[0])
	srcUrl := html.EscapeString(parts[1])

	return "<img src=\"" + srcUrl + "\" alt=\"" + altText + "\" class=\"max-w-full h-auto rounded-lg my-4\">"
}

func (r *Renderer) linkRenderer(value string) string {
	parts := strings.Split(value, " -> ")
	if len(parts) != 2 {
		return "<a href=\"\">Invalid link format</a>"
	}

	linkText := html.EscapeString(parts[0])
	hrefUrl := html.EscapeString(parts[1])

	return "<a href=\"" + hrefUrl + "\" class=\"text-blue-600 hover:underline\">" + linkText + "</a>"
}
