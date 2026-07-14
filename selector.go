package htmlbuilder

import (
	"fmt"
	"strings"
	"sync/atomic"
)

var classCounter atomic.Uint64

// selectorRule holds one generated CSS selector rule, e.g.
// ".hb-1:hover{...}" or ".hb-2::before{...}".
type selectorRule struct {
	class    string
	selector string
	styles   []attr
}

// StyleClass generates a unique class for these styles instead of writing them inline.
// This allows pseudo-classes like :hover to override the base properties naturally.
func (n *Node) StyleClass(s Style) *Node {
	return n.addRule("", s) // Empty string selector creates a base class rule (e.g. ".hb-1 { ... }")
}

// Hover attaches a :hover rule.
func (n *Node) Hover(s Style) *Node {
	return n.addRule(":hover", s)
}

// Focus attaches a :focus rule.
func (n *Node) Focus(s Style) *Node {
	return n.addRule(":focus", s)
}

// Active attaches an :active rule.
func (n *Node) Active(s Style) *Node {
	return n.addRule(":active", s)
}

// Before attaches a ::before rule.
func (n *Node) Before(s Style) *Node {
	return n.addRule("::before", s)
}

// After attaches a ::after rule.
func (n *Node) After(s Style) *Node {
	return n.addRule("::after", s)
}

// Pseudo attaches any pseudo-element, e.g. "before", "after", "selection".
func (n *Node) Pseudo(name string, s Style) *Node {
	return n.addRule("::"+name, s)
}

// HoverBefore attaches a :hover::before rule.
func (n *Node) HoverBefore(s Style) *Node {
	return n.addRule(":hover::before", s)
}

// HoverAfter attaches a :hover::after rule.
func (n *Node) HoverAfter(s Style) *Node {
	return n.addRule(":hover::after", s)
}

// HoverPseudo attaches a :hover::<name> rule.
func (n *Node) HoverPseudo(name string, s Style) *Node {
	return n.addRule(":hover::"+name, s)
}

func (n *Node) addRule(selector string, s Style) *Node {
	class := fmt.Sprintf("hb-%d", classCounter.Add(1))
	n.Class(class)

	n.selectorRules = append(n.selectorRules, selectorRule{
		class:    class,
		selector: selector,
		styles:   styleToAttrs(s),
	})

	return n
}

// collectSelectorRules walks the subtree and gathers generated selector rules.
func (n *Node) collectSelectorRules() []selectorRule {
	var rules []selectorRule

	rules = append(rules, n.selectorRules...)

	for _, c := range n.children {
		rules = append(rules, c.collectSelectorRules()...)
	}

	return rules
}

// renderSelectorRuleCSS renders all generated selector rules into CSS.
func renderSelectorRuleCSS(rules []selectorRule) string {
	var sb strings.Builder

	for _, r := range rules {
		sb.WriteString(".")
		sb.WriteString(r.class)
		sb.WriteString(r.selector)
		sb.WriteString("{")

		for _, s := range r.styles {
			sb.WriteString(s.key)
			sb.WriteString(":")
			sb.WriteString(s.value)
			sb.WriteString(" !important;") // overrides inline styles!
		}

		sb.WriteString("}")
	}

	return sb.String()
}
