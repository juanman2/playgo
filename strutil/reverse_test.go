package strutil

import (
	"testing"
)

// TestReverse verifies the function reverse
func TestReverse(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"1234", "4321"},
	}

	for _, tt := range tests {
		if got := Reverse(tt.in); got != tt.out {
			t.Errorf("Reverse(%v) want %v got %v", tt.in, tt.out, got)
		}
	}
}
