package htmlbuilder

import "testing"

func TestGeneratedDiv_WrapsChildren(t *testing.T) {
	d := Div(Span("a"), Span("b"))
	if d.tag != "div" || len(d.children) != 2 {
		t.Fatalf("expected div with 2 children, got tag=%s children=%d", d.tag, len(d.children))
	}
}

func TestGeneratedSpan_SetsText(t *testing.T) {
	s := Span("hello")
	if s.tag != "span" || s.text != "hello" {
		t.Fatalf("unexpected node: %+v", s)
	}
}

func TestGeneratedA_RequiresAttrForHref(t *testing.T) {
	link := A("Home").Attr("href", "/")
	if link.tag != "a" || link.text != "Home" {
		t.Fatalf("unexpected node: %+v", link)
	}
	if len(link.attrs) != 1 || link.attrs[0].key != "href" || link.attrs[0].value != "/" {
		t.Fatalf("expected href=/ attr set via .Attr(), got %+v", link.attrs)
	}
}

func TestGeneratedImg_IsVoidNoArgs(t *testing.T) {
	img := Img().Attr("src", "/logo.png").Attr("alt", "logo")
	if img.tag != "img" {
		t.Fatalf("expected tag img, got %s", img.tag)
	}
	if !voidElements["img"] {
		t.Fatal("img should be registered as void")
	}
}

func TestGeneratedBr_IsVoid(t *testing.T) {
	b := Br()
	if b.tag != "br" || !voidElements["br"] {
		t.Fatalf("expected br to be a void element, got %+v", b)
	}
}

func TestGeneratedUl_Li_Nesting(t *testing.T) {
	list := Ul(Li(Span("item one")), Li(Span("item two")))
	if list.tag != "ul" || len(list.children) != 2 {
		t.Fatalf("expected ul with 2 li children, got %+v", list)
	}
}
