package main

import (
	"fmt"

	"github.com/Thruqe/htmlbuilder"
)

// onClick wires an element's click to POST to url and swap the
// response HTML into the element with id targetID.
func onClick(n *htmlbuilder.Node, url, targetID string) *htmlbuilder.Node {
	return n.Attr("onclick", fmt.Sprintf("__hb_call(%q,%q)", url, targetID))
}

// themeToggle marks a node as a theme-toggle button. iconSelector
// points at the icon element inside it to swap classes on click.
func themeToggle(n *htmlbuilder.Node, iconSelector string) *htmlbuilder.Node {
	return n.Attr("onclick", fmt.Sprintf("__hb_toggle_theme(%q)", iconSelector))
}
