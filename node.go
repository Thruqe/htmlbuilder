package htmlbuilder

// voidElements are HTML elements that never have closing tags or children.
var voidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true,
	"embed": true, "hr": true, "img": true, "input": true,
	"link": true, "meta": true, "param": true, "source": true,
	"track": true, "wbr": true,
}

// attr is a single ordered key/value pair, used instead of a map
// so render output is deterministic.
type attr struct {
	key   string
	value string
}

type Node struct {
	tag      string
	attrs    []attr
	classes  []string
	styles   []attr
	children []*Node
	text     string
	rawHTML  string
	parent   *Node

	// Generated CSS selector rules (e.g. :hover, ::before, :hover::before).
	selectorRules []selectorRule
}

// El creates a new element node with the given tag name.
func El(tag string) *Node {
	return &Node{tag: tag}
}

// Attr sets an arbitrary HTML attribute. Repeated calls with the
// same key overwrite the previous value.
func (n *Node) Attr(key, value string) *Node {
	for i, a := range n.attrs {
		if a.key == key {
			n.attrs[i].value = value
			return n
		}
	}
	n.attrs = append(n.attrs, attr{key, value})
	return n
}

// Class appends one or more CSS classes.
func (n *Node) Class(classes ...string) *Node {
	n.classes = append(n.classes, classes...)
	return n
}

// Text sets escaped text content for this node.
func (n *Node) Text(t string) *Node {
	n.text = t
	return n
}

// Raw sets unescaped HTML content for this node. Caller is responsible
// for ensuring the content is safe — no escaping is applied at render time.
func (n *Node) Raw(html string) *Node {
	n.rawHTML = html
	return n
}

// Child appends one or more child nodes. Ignored for void elements.
func (n *Node) Child(children ...*Node) *Node {
	for _, c := range children {
		c.parent = n
	}
	n.children = append(n.children, children...)
	return n
}

// If conditionally applies fn to the node, returning n either way —
// keeps the fluent chain intact without needing an if-statement break.
func (n *Node) If(cond bool, fn func(*Node)) *Node {
	if cond {
		fn(n)
	}
	return n
}

// Each runs fn once per item in items, returning one node per item.
// Useful in place of {{range}}: Ul(Each(items, func(i Item) *Node {...})...)
func Each[T any](items []T, fn func(T) *Node) []*Node {
	nodes := make([]*Node, 0, len(items))
	for _, item := range items {
		nodes = append(nodes, fn(item))
	}
	return nodes
}
