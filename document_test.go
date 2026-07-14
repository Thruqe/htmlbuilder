package htmlbuilder

import (
	"strings"
	"testing"
)

func TestNew_DefaultsToEnglishLang(t *testing.T) {
	d := New()
	if d.lang != "en" {
		t.Fatalf("expected default lang 'en', got %q", d.lang)
	}
}

func TestLang_Overrides(t *testing.T) {
	d := New().Lang("fr")
	if d.lang != "fr" {
		t.Fatalf("expected lang 'fr', got %q", d.lang)
	}
}

func TestTitle_AddsTitleTag(t *testing.T) {
	d := New().Title("My Page")
	if len(d.head.children) != 1 || d.head.children[0].tag != "title" {
		t.Fatalf("expected single title child in head, got %+v", d.head.children)
	}
	if d.head.children[0].text != "My Page" {
		t.Fatalf("expected title text 'My Page', got %q", d.head.children[0].text)
	}
}

func TestTitle_ReplacesRatherThanAppends(t *testing.T) {
	d := New().Title("First").Title("Second")
	count := 0
	for _, c := range d.head.children {
		if c.tag == "title" {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("expected exactly 1 title tag, got %d", count)
	}
	if d.head.children[0].text != "Second" {
		t.Fatalf("expected title replaced with 'Second', got %q", d.head.children[0].text)
	}
}

func TestMeta_AppendsToHead(t *testing.T) {
	d := New().Meta(map[string]string{"charset": "utf-8"})
	if len(d.head.children) != 1 || d.head.children[0].tag != "meta" {
		t.Fatalf("expected meta tag in head, got %+v", d.head.children)
	}
}

func TestLink_AppendsToHead(t *testing.T) {
	d := New().Link(map[string]string{"rel": "stylesheet", "href": "/style.css"})
	if len(d.head.children) != 1 || d.head.children[0].tag != "link" {
		t.Fatalf("expected link tag in head, got %+v", d.head.children)
	}
}

func TestBody_ReturnsAssignableNode(t *testing.T) {
	d := New()
	d.Body().Child(El("h1").Text("hi"))
	if len(d.body.children) != 1 {
		t.Fatalf("expected 1 child in body, got %d", len(d.body.children))
	}
}

func TestDocument_RenderIncludesDoctypeAndLang(t *testing.T) {
	d := New().Title("Test")
	out := d.String()

	if !strings.HasPrefix(out, "<!DOCTYPE html>") {
		t.Fatalf("expected output to start with doctype, got: %s", out)
	}
	if !strings.Contains(out, `<html lang="en">`) {
		t.Fatalf("expected html lang attribute, got: %s", out)
	}
	if !strings.Contains(out, "<title>Test</title>") {
		t.Fatalf("expected title tag rendered, got: %s", out)
	}
	if !strings.HasSuffix(out, "</html>") {
		t.Fatalf("expected output to end with </html>, got: %s", out)
	}
}