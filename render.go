package htmlbuilder

import (
	"fmt"
	"html"
	"io"
	"strings"
)

// Render writes the Document as a complete HTML document, including
// doctype, to w.
func (d *Document) Render(w io.Writer) {
	rules := d.body.collectSelectorRules()
	if len(rules) > 0 {
		css := renderSelectorRuleCSS(rules)
		d.head.Child(El("style").Raw(css))
	}

	fmt.Fprint(w, "<!DOCTYPE html>")
	fmt.Fprintf(w, `<html lang="%s">`, html.EscapeString(d.lang))
	d.head.Render(w)
	d.body.Render(w)
	fmt.Fprint(w, "</html>")
}

// String renders the Document to a string. Convenience wrapper around
// Render for cases where you don't have an io.Writer handy (tests,
// debugging, etc.) — prefer Render(w) in HTTP handlers to avoid the
// extra allocation.
func (d *Document) String() string {
	var sb strings.Builder
	d.Render(&sb)
	return sb.String()
}

// Render writes the Node (and its full subtree) as HTML to w.
func (n *Node) Render(w io.Writer) {
	fmt.Fprintf(w, "<%s", n.tag)
	n.renderAttrs(w)
	n.renderClasses(w)
	n.renderStyles(w)

	if voidElements[n.tag] {
		fmt.Fprint(w, ">")
		return
	}
	fmt.Fprint(w, ">")

	if n.rawHTML != "" {
		fmt.Fprint(w, n.rawHTML)
	} else if n.text != "" {
		fmt.Fprint(w, html.EscapeString(n.text))
	}

	for _, child := range n.children {
		child.Render(w)
	}

	fmt.Fprintf(w, "</%s>", n.tag)
}

// String renders the Node (and its subtree) to a string.
func (n *Node) String() string {
	var sb strings.Builder
	n.Render(&sb)
	return sb.String()
}

func (n *Node) renderAttrs(w io.Writer) {
	for _, a := range n.attrs {
		fmt.Fprintf(w, ` %s="%s"`, a.key, html.EscapeString(a.value))
	}
}

func (n *Node) renderClasses(w io.Writer) {
	if len(n.classes) == 0 {
		return
	}
	fmt.Fprintf(w, ` class="%s"`, html.EscapeString(strings.Join(n.classes, " ")))
}

func (n *Node) renderStyles(w io.Writer) {
	if len(n.styles) == 0 {
		return
	}
	var sb strings.Builder
	for _, s := range n.styles {
		sb.WriteString(s.key)
		sb.WriteString(":")
		sb.WriteString(s.value)
		sb.WriteString(";")
	}
	fmt.Fprintf(w, ` style="%s"`, html.EscapeString(sb.String()))
}
