package main

import (
	"strings"
	"testing"
)

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple property",
			in:   "color",
			want: "Color",
		},
		{
			name: "kebab case",
			in:   "background-color",
			want: "BackgroundColor",
		},
		{
			name: "multiple words",
			in:   "border-top-left-radius",
			want: "BorderTopLeftRadius",
		},
		{
			name: "vendor prefix",
			in:   "-webkit-appearance",
			want: "WebkitAppearance",
		},
		{
			name: "leading multiple dashes",
			in:   "--custom-property",
			want: "CustomProperty",
		},
		{
			name: "empty string",
			in:   "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toPascalCase(tt.in)

			if got != tt.want {
				t.Errorf(
					"toPascalCase(%q) = %q, want %q",
					tt.in,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestToPascalCaseProducesValidGoIdentifiers(t *testing.T) {
	properties := []string{
		"background-color",
		"font-size",
		"border-radius",
		"-webkit-transform",
		"animation-name",
	}

	for _, prop := range properties {
		got := toPascalCase(prop)

		if got == "" {
			t.Errorf("expected identifier for %q", prop)
		}

		if strings.ContainsAny(got, "- ") {
			t.Errorf(
				"generated identifier %q contains invalid characters",
				got,
			)
		}

		if got[0] < 'A' || got[0] > 'Z' {
			t.Errorf(
				"generated identifier %q is not PascalCase",
				got,
			)
		}
	}
}

func TestToPascalCaseDoesNotLeaveEmptySegments(t *testing.T) {
	tests := []string{
		"-webkit-user-select",
		"--foo",
		"---bar",
	}

	for _, input := range tests {
		got := toPascalCase(input)

		if strings.Contains(got, "-") {
			t.Errorf(
				"%q produced invalid identifier %q",
				input,
				got,
			)
		}
	}
}

func TestCSSPropertyFilteringRules(t *testing.T) {
	properties := []string{
		"color",
		"background-color",
		"--custom-property",
		"--*",
	}

	var names []string

	for _, prop := range properties {
		if strings.HasPrefix(prop, "--") {
			continue
		}

		names = append(names, prop)
	}

	if len(names) != 2 {
		t.Fatalf(
			"expected 2 CSS properties after filtering, got %d",
			len(names),
		)
	}

	if names[0] != "color" || names[1] != "background-color" {
		t.Errorf("unexpected filtered properties: %#v", names)
	}
}

func TestGeneratedFieldNamesSkipCollisions(t *testing.T) {
	properties := []string{
		"background-color",
		"background--color",
		"border-radius",
	}

	seen := make(map[string]bool)
	generated := 0

	for _, prop := range properties {
		field := toPascalCase(prop)

		if seen[field] {
			continue
		}

		seen[field] = true
		generated++
	}

	if generated != 2 {
		t.Fatalf(
			"expected 2 generated fields after collision filtering, got %d",
			generated,
		)
	}
}
