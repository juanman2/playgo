package strutil

import "testing"

func TestNumTilePossibilities(t *testing.T) {

	tests := []struct {
		in  string
		out int
	}{
		{"AAB", 8},
		{"ABC", 12},
		{"ABCD", 30},
		{"AAAA", 4},
		{"AAAC", 13},
		{"AAABBC", 188},
	}

	for _, tt := range tests {
		if count := numTilePossibilities(tt.in); count != tt.out {
			t.Errorf("numTilePossibilities(%s) returned %d but expected %d\n",
				tt.in, count, tt.out)
		}
	}
}
