package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		in, sep string
		want    []string
	}{
		{"", "", []string{}},
		{"", ":", []string{""}},
		{":", ":", []string{"", ""}},
		{"abc", ":", []string{"abc"}},
		{"abc", "", []string{"a", "b", "c"}},
		{":a", ":", []string{"", "a"}},
		{"a:", ":", []string{"a", ""}},
		{"a:bc", ":", []string{"a", "bc"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a::b::c", ":", []string{"a", "", "b", "", "c"}},
		{"a::b:c", "::", []string{"a", "b:c"}},
	}

	for _, test := range tests {
		words := strings.Split(test.in, test.sep)
		if len(words) != len(test.want) {
			t.Errorf("Split(%q, %q) returned %v (len %d), want %v (len %d)",
				test.in, test.sep, words, len(words), test.want, len(test.want))
			continue
		}
		for i, word := range words {
			if word != test.want[i] {
				t.Errorf("Split(%q, %q) returned %v, want %v", test.in, test.sep, words, test.want)
			}
		}
	}

}
