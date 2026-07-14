package htmlbuilder

import "testing"

func TestPx_Int(t *testing.T) {
	if px(16) != "16px" {
		t.Errorf("px(16) = %q, want 16px", px(16))
	}
}

func TestPx_Float(t *testing.T) {
	if px(1.5) != "1.5px" {
		t.Errorf("px(1.5) = %q, want 1.5px", px(1.5))
	}
}

func TestPct(t *testing.T) {
	if pct(50) != "50%" {
		t.Errorf("pct(50) = %q, want 50%%", pct(50))
	}
}

func TestRem(t *testing.T) {
	if rem(1.5) != "1.5rem" {
		t.Errorf("rem(1.5) = %q, want 1.5rem", rem(1.5))
	}
}