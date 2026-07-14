// style.go
package htmlbuilder

// SetStyle sets one raw CSS property directly — escape hatch for
// properties not covered by the generated Style struct (e.g. an
// unreleased/experimental property, or a custom property like
// "--my-var").
func (n *Node) SetStyle(prop, value string) *Node {
	n.setStyle(prop, value)
	return n
}

// setStyle is the internal overwrite-aware setter both CSS()
// (in style_generated.go) and SetStyle route through.
func (n *Node) setStyle(prop, value string) {
	for i, s := range n.styles {
		if s.key == prop {
			n.styles[i].value = value
			return
		}
	}
	n.styles = append(n.styles, attr{prop, value})
}