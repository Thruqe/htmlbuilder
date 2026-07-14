package htmlbuilder

import "testing"

func TestAttr_SetsValue(t *testing.T) {
	n := El("div").Attr("id", "main")
	if len(n.attrs) != 1 || n.attrs[0].key != "id" || n.attrs[0].value != "main" {
		t.Fatalf("expected attr id=main, got %+v", n.attrs)
	}
}

func TestAttr_OverwritesExistingKey(t *testing.T) {
	n := El("div").Attr("id", "first").Attr("id", "second")
	if len(n.attrs) != 1 {
		t.Fatalf("expected 1 attr after overwrite, got %d", len(n.attrs))
	}
	if n.attrs[0].value != "second" {
		t.Fatalf("expected overwritten value 'second', got %q", n.attrs[0].value)
	}
}

func TestClass_Appends(t *testing.T) {
	n := El("div").Class("card").Class("shadow", "rounded")
	want := []string{"card", "shadow", "rounded"}
	if len(n.classes) != len(want) {
		t.Fatalf("expected %d classes, got %d", len(want), len(n.classes))
	}
	for i, c := range want {
		if n.classes[i] != c {
			t.Errorf("class[%d] = %q, want %q", i, n.classes[i], c)
		}
	}
}

func TestText_SetsContent(t *testing.T) {
	n := El("p").Text("hello")
	if n.text != "hello" {
		t.Fatalf("expected text 'hello', got %q", n.text)
	}
}

func TestRaw_SetsUnescapedContent(t *testing.T) {
	n := El("div").Raw("<b>bold</b>")
	if n.rawHTML != "<b>bold</b>" {
		t.Fatalf("expected rawHTML set, got %q", n.rawHTML)
	}
}

func TestChild_SetsParent(t *testing.T) {
	parent := El("div")
	child := El("span")
	parent.Child(child)

	if child.parent != parent {
		t.Fatal("expected child.parent to point to parent node")
	}
	if len(parent.children) != 1 || parent.children[0] != child {
		t.Fatal("expected parent.children to contain child")
	}
}

func TestChild_MultipleChildrenPreserveOrder(t *testing.T) {
	parent := El("ul").Child(El("li").Text("a"), El("li").Text("b"), El("li").Text("c"))
	if len(parent.children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(parent.children))
	}
	want := []string{"a", "b", "c"}
	for i, w := range want {
		if parent.children[i].text != w {
			t.Errorf("children[%d].text = %q, want %q", i, parent.children[i].text, w)
		}
	}
}

func TestIf_AppliesWhenTrue(t *testing.T) {
	n := El("div").If(true, func(n *Node) {
		n.Class("active")
	})
	if len(n.classes) != 1 || n.classes[0] != "active" {
		t.Fatalf("expected class 'active' applied, got %+v", n.classes)
	}
}

func TestIf_SkipsWhenFalse(t *testing.T) {
	n := El("div").If(false, func(n *Node) {
		n.Class("active")
	})
	if len(n.classes) != 0 {
		t.Fatalf("expected no classes, got %+v", n.classes)
	}
}

func TestEach_BuildsNodesFromItems(t *testing.T) {
	items := []string{"a", "b", "c"}
	nodes := Each(items, func(s string) *Node {
		return El("li").Text(s)
	})

	if len(nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(nodes))
	}
	for i, s := range items {
		if nodes[i].text != s {
			t.Errorf("node[%d].text = %q, want %q", i, nodes[i].text, s)
		}
	}
}

func TestEach_UsableAsChildren(t *testing.T) {
	items := []int{1, 2, 3}
	ul := El("ul").Child(Each(items, func(i int) *Node {
		return El("li")
	})...)

	if len(ul.children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(ul.children))
	}
}

func TestVoidElements_ContainsExpected(t *testing.T) {
	expected := []string{"br", "img", "input", "meta", "hr"}
	for _, tag := range expected {
		if !voidElements[tag] {
			t.Errorf("expected %q to be a void element", tag)
		}
	}
	if voidElements["div"] {
		t.Error("div should not be a void element")
	}
}