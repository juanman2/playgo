// Copyright 2020 Juan Tellez All rights reserved.

package strutil

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"hello", false},
		{"abbac", false},
		{"abba", true},
	}

	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.out {
			t.Errorf("IsPalindrome(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"hello", ""},
		{"abbac", "abba"},
		{"papa", "apa"},
	}

	for _, tt := range tests {
		if got := ValidPalindrome(tt.in); got != tt.out {
			t.Errorf("ValidPalindrome(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}

}

func TestCouldBePalindrome(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"hello", false},
		{"abbac", true},
		{"papa", true},
		{"123456789009876542", false},
		{"papaxapap", true},
		{"xpapaxapap", true},
		{"papaxapapx", true},
	}

	for _, tt := range tests {
		if got := CouldBePalindrome(tt.in); got != tt.out {
			t.Errorf("CouldBePalindrome(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}
}

func TestTrimIth(t *testing.T) {

	tests := []struct {
		in  string
		i   int
		out string
	}{
		{"hello", 3, "helo"},
		{"abbac", 0, "bbac"},
		{"papa", 3, "pap"},
	}

	for _, tt := range tests {
		if got := TrimIth(tt.in, tt.i); got != tt.out {
			t.Errorf("TrimIth(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}
}
