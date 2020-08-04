package overlapts

import "testing"

func TestOverlapTimes(t *testing.T) {
	tests := []struct {
		in  []td
		out int
	}{
		{[]td{{1, 3}, {1, 2}, {4, 5}}, 1},
		{[]td{{1, 3}, {1, 2}, {5, 5}}, 2},
		{[]td{{1, 1}, {3, 1}, {5, 1}}, 3},
		{[]td{{1, 1}, {3, 1}, {5, 1}}, 3},
	}

	for _, tt := range tests {
		c := overlapTimes(tt.in)
		if c != tt.out {
			t.Errorf("Called ot(%v) got %d but expected %d", tt.in, c, tt.out)
		}
	}
}
