package attendance

import "testing"

func TestCheckRecord(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
	}

	for _, tt := range tests {
		passed := checkRecord(tt.in)
		if passed != tt.out {
			t.Errorf("checkRecord failed(%s) returned %v expected %v", tt.in, passed, tt.out)
		}
	}
}
