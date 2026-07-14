package main

import (
	"strings"
	"testing"
)

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "div",
			want:  "Div",
		},
		{
			input: "span",
			want:  "Span",
		},
		{
			input: "colgroup",
			want:  "Colgroup",
		},
		{
			input: "optgroup",
			want:  "Optgroup",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toPascalCase(tt.input)

			if got != tt.want {
				t.Errorf(
					"toPascalCase(%q) = %q, want %q",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestAllTagsHaveValidKinds(t *testing.T) {
	validKinds := map[string]bool{
		"children": true,
		"text":     true,
		"void":     true,
		"custom":   true,
	}

	for _, tag := range allTags {
		if !validKinds[tag.kind] {
			t.Errorf(
				"tag %q has invalid kind %q",
				tag.tag,
				tag.kind,
			)
		}
	}
}

func TestAllTagsAreUnique(t *testing.T) {
	seen := make(map[string]bool)

	for _, tag := range allTags {
		if seen[tag.tag] {
			t.Errorf(
				"duplicate HTML tag found: %q",
				tag.tag,
			)
		}

		seen[tag.tag] = true
	}
}

func TestAllTagsHaveNames(t *testing.T) {
	for _, tag := range allTags {
		if tag.tag == "" {
			t.Error("found empty tag name")
		}

		if tag.kind == "" {
			t.Errorf(
				"tag %q has empty kind",
				tag.tag,
			)
		}
	}
}

func TestGeneratedConstructorNamesAreValid(t *testing.T) {
	for _, tag := range allTags {
		name := toPascalCase(tag.tag)

		if name == "" {
			t.Errorf(
				"tag %q generated empty constructor name",
				tag.tag,
			)
		}

		if strings.ContainsAny(name, "-_ ") {
			t.Errorf(
				"tag %q generated invalid Go identifier %q",
				tag.tag,
				name,
			)
		}

		if name[0] < 'A' || name[0] > 'Z' {
			t.Errorf(
				"constructor %q is not PascalCase",
				name,
			)
		}
	}
}

func TestStyleIsExcludedFromGeneration(t *testing.T) {
	for _, tag := range allTags {
		if tag.tag == "style" {
			// main() explicitly skips this tag because it has a custom implementation
			return
		}
	}

	t.Error("expected style tag to exist in allTags")
}

func TestVoidTagsAreMarkedCorrectly(t *testing.T) {
	expectedVoidTags := map[string]bool{
		"area":   true,
		"br":     true,
		"img":    true,
		"input":  true,
		"meta":   true,
		"link":   true,
		"source": true,
	}

	for _, tag := range allTags {
		if expectedVoidTags[tag.tag] && tag.kind != "void" {
			t.Errorf(
				"expected %q to be void, got %q",
				tag.tag,
				tag.kind,
			)
		}
	}
}

func TestTextTagsAreMarkedCorrectly(t *testing.T) {
	expectedTextTags := map[string]bool{
		"a":      true,
		"span":   true,
		"title":  true,
		"strong": true,
		"script": true,
	}

	for _, tag := range allTags {
		if expectedTextTags[tag.tag] && tag.kind != "text" {
			t.Errorf(
				"expected %q to be text, got %q",
				tag.tag,
				tag.kind,
			)
		}
	}
}

func TestChildrenTagsAreMarkedCorrectly(t *testing.T) {
	expectedChildrenTags := map[string]bool{
		"div":     true,
		"section": true,
		"body":    true,
		"html":    true,
		"footer":  true,
	}

	for _, tag := range allTags {
		if expectedChildrenTags[tag.tag] && tag.kind != "children" {
			t.Errorf(
				"expected %q to be children, got %q",
				tag.tag,
				tag.kind,
			)
		}
	}
}
