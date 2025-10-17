package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q, expected %q", repeated, expected)

	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeat := Repeat("b")
	fmt.Println(repeat)
	// Output: bbbbb
}
