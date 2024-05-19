package injiection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Longchen")

	got := buf.String()
	want := "Hello, Longchen"

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
