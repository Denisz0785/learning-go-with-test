package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	got := Repeat("a", 5)

	newFunction(expected, got, t)
}

func newFunction(expected string, got string, t *testing.T) {
	if expected != got {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 5)
	fmt.Println(got)
	// Output: aaaaa
}
