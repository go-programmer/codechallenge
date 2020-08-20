package wordcount

import (
	"fmt"
	"testing"
)

func TestValidWord(t *testing.T) {
	someStr := " a  this is some string a a a"

	want := make(map[string]int)
	want["a"] = 4
	got := wordCount(someStr)

	fmt.Println(got)

	if got["a"] != want["a"] {
		t.Errorf("Hello() = %q, want %q", got["a"], want["a"])
	}
}
