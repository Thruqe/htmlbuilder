package htmlbuilder

import "testing"

func TestCSS_SkipsZeroValueFields(t *testing.T) {
	n := Div().CSS(Style{Padding: "1rem"})
	if len(n.styles) != 1 {
		t.Fatalf("expected 1 style set, got %d: %+v", len(n.styles), n.styles)
	}
	if n.styles[0].key != "padding" || n.styles[0].value != "1rem" {
		t.Fatalf("expected padding=1rem, got %+v", n.styles[0])
	}
}

func TestCSS_SetsMultipleFields(t *testing.T) {
	n := Div().CSS(Style{Padding: "1rem", Color: "#333", Width: "100%"})
	if len(n.styles) != 3 {
		t.Fatalf("expected 3 styles, got %d: %+v", len(n.styles), n.styles)
	}
}

func TestCSS_OverwritesOnRepeatedCall(t *testing.T) {
	n := Div().CSS(Style{Padding: "1rem"}).CSS(Style{Padding: "2rem"})
	if len(n.styles) != 1 || n.styles[0].value != "2rem" {
		t.Fatalf("expected single overwritten padding=2rem, got %+v", n.styles)
	}
}

func TestSetStyle_RawEscapeHatch(t *testing.T) {
	n := Div().SetStyle("z-index", "10")
	if len(n.styles) != 1 || n.styles[0].key != "z-index" {
		t.Fatalf("expected z-index style set, got %+v", n.styles)
	}
}
