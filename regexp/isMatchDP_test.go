package playgo

import (
	"fmt"
	"testing"
)

func TestIsRegexpMatch(t *testing.T) {
	tests := []struct {
		inS string
		inP string
		out bool
	}{
		{"aaca", ".", false},
		{"aacb", ".*", true},
		{"aacc", "a*.c", true},
		{"aa", "a", false},
		{"mississippi", "mis*is*ip*.", true},
		{"aaa", "a*a", true},
		{"ab", ".*c", false},
	}

	for _, tt := range tests {
		match := IsRegexpMatchDP(tt.inS, tt.inP)
		fmt.Printf("IsRegexpMatchDP(%s, %s) returned %v expected %v\n",
			tt.inS, tt.inP, match, tt.out)
		if match != tt.out {
			t.Errorf("IsRegexpMatch(%s, %s) returned %v expected %v",
				tt.inS, tt.inP, match, tt.out)
		}
	}
}
