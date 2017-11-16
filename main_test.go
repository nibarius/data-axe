package dataaxe

import (
	"testing"
)

func TestSanitize(t *testing.T) {
	in := []string{"<br>a</br>", "<b>b</b>"}
	out := sanitize(in)

	if out[0] != "a" {
		t.Errorf("Sanitize failed expected %s, got %s", "a", out[0])
	}
}

func TestSanitizeEmpty(t *testing.T) {
	in := make([]string, 0)
	out := sanitize(in)

	if len(out) != 0 {
		t.Errorf("Sanitize failed empty input does not give empty output")
	}
}

func TestSanitizeNil(t *testing.T) {
	out := sanitize(nil)

	if out != nil {
		t.Errorf("Sanitize failed nil input does not give nil output")
	}
}
