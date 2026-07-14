package htmlbuilder

import (
	"strings"
	"testing"
)

func TestRender_SimpleTextNode(t *testing.T) {
	out := El("p").Text("hello").String()
	if out != "<p>hello</p>" {
		t.Fatalf("got %q", out)
	}
}

func TestRender_EscapesTextContent(t *testing.T) {
	out := El("p").Text("<script>alert(1)</script>").String()
	if strings.Contains(out, "<script>") {
		t.Fatalf("expected text to be escaped, got %q", out)
	}
	if !strings.Contains(out, "&lt;script&gt;") {
		t.Fatalf("expected escaped script tag, got %q", out)
	}
}

func TestRender_RawSkipsEscaping(t *testing.T) {
	out := El("div").Raw("<b>bold</b>").String()
	if out != "<div><b>bold</b></div>" {
		t.Fatalf("expected raw HTML preserved, got %q", out)
	}
}

func TestRender_EscapesAttrValues(t *testing.T) {
	out := El("a").Attr("href", `"><script>`).String()
	if strings.Contains(out, "<script>") {
		t.Fatalf("expected attr value to be escaped, got %q", out)
	}
}

func TestRender_ClassesJoinedWithSpace(t *testing.T) {
	out := El("div").Class("card", "shadow").String()
	if !strings.Contains(out, `class="card shadow"`) {
		t.Fatalf(`expected class="card shadow", got %q`, out)
	}
}

func TestRender_StylesRenderedInOrder(t *testing.T) {
	node := El("div").
		SetStyle("color", "#333").
		SetStyle("padding", "1rem")

	got := node.String()

	expected := `style="color: #333; padding: 1rem;"`

	if !strings.Contains(got, expected) {
		t.Fatalf(
			"expected ordered inline style, got %q",
			got,
		)
	}
}

func TestRender_VoidElementNoClosingTag(t *testing.T) {
	out := Br().String()
	if out != "<br>" {
		t.Fatalf(`expected "<br>", got %q`, out)
	}
}

func TestRender_ImgVoidElementWithAttrs(t *testing.T) {
	out := Img().Attr("src", "/logo.png").Attr("alt", "logo").String()
	if !strings.HasPrefix(out, "<img ") || strings.Contains(out, "</img>") {
		t.Fatalf("expected self-contained void img tag, got %q", out)
	}
}

func TestRender_NestedChildren(t *testing.T) {
	out := Div(H1("Welcome"), P("intro")).String()
	want := "<div><h1>Welcome</h1><p>intro</p></div>"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestRender_DeepNesting(t *testing.T) {
	out := Ul(Li(Span("item one"))).String()
	want := "<ul><li><span>item one</span></li></ul>"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}
