package playgo

import (
	"fmt"
	"testing"
)

func TestDotMatch(t *testing.T) {

	tests := []struct {
		in     string
		inIdx  int
		outIdx int
		outB   int
	}{
		{"a", 0, 1, beginState},
		{"a", 1, 1, errorState},
		{"ai", 1, 2, beginState},
	}

	for _, tt := range tests {
		gotB, gotI := DotMatch(tt.in, tt.inIdx)
		if gotB != tt.outB || gotI != tt.outIdx {
			t.Errorf(
				"DotMatch(%v %v) got %v %v but expected %v %v",
				tt.in, tt.inIdx, gotB, gotI, tt.outB, tt.outIdx)
		}
	}

}

func TestStarMatch(t *testing.T) {
	tests := []struct {
		in     string
		inP    byte
		inIdx  int
		outIdx int
		outB   int
	}{
		{"a", 'a', 0, 1, beginState},
		{"aa", 'a', 0, 2, beginState},
		{"aca", 'a', 1, 1, beginState},
		{"aca", 'c', 1, 2, beginState},
		{"aca", '.', 0, 3, beginState},
		{"aaca", '.', 0, 4, beginState},
		{"", '.', 0, 0, beginState},
	}

	for _, tt := range tests {
		gotB, gotI := StarMatch(tt.in, tt.inP, tt.inIdx)
		if gotB != tt.outB || gotI != tt.outIdx {
			t.Errorf(
				"StarMatch(%v %v %v) got %v %v but expected %v %v",
				tt.in, tt.inP, tt.inIdx, gotB, gotI, tt.outB, tt.outIdx)
		}
	}
}

func TestRegexpState(t *testing.T) {
	tests := []struct {
		inP      string
		inIdx    int
		outState int
		outIdx   int
		outCh    byte
	}{
		{"", 0, endState, 0, 0},
		{".", 0, dotState, 1, 0},
		{"a*", 0, starState, 2, 'a'},
		{"a*", 2, endState, 2, 0},
		{"x", 0, simpleMatchState, 1, 'x'},
		{"mi", 1, simpleMatchState, 2, 'i'},
	}

	for _, tt := range tests {
		state, idx, ch := RegexpState(tt.inP, tt.inIdx)
		if state != tt.outState || idx != tt.outIdx {
			t.Errorf("RegexpState(%v , %v ) got state:%d idx:%d ch:%c expected %d %d %c",
				tt.inP, tt.inIdx,
				state, idx, ch,
				tt.outState, tt.outIdx, tt.outCh)
		}
	}
}

func TestSimpleMatch(t *testing.T) {
	tests := []struct {
		inS      string
		inP      byte
		inIdx    int
		outState int
		outIdx   int
	}{
		{"aaca", 'c', 0, errorState, 0},
		{"aaca", 'a', 0, beginState, 1},
		{"aaca", 'c', 2, beginState, 3},
		{"aa", 'a', 2, endState, 2},
		{"a", 'a', 0, beginState, 1},
		{"miss", 'm', 0, beginState, 1},
		{"miss", 'i', 1, beginState, 2},
		{"miss", 's', 2, beginState, 3},
		{"miss", 's', 3, beginState, 4},
	}

	for _, tt := range tests {
		state, idx := SimpleMatch(tt.inS, tt.inP, tt.inIdx)
		if state != tt.outState || idx != tt.outIdx {
			t.Errorf(
				"SimpleMatch(%s, %c, %d) returned %d, %d expected %d, %d",
				tt.inS, tt.inP, tt.inIdx,
				state, idx,
				tt.outState, tt.outIdx)
		}
	}
}

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
		match := IsRegexpMatch(tt.inS, tt.inP)
		fmt.Printf("IsRegexpMatch(%s, %s) returned %v expected %v\n",
			tt.inS, tt.inP, match, tt.out)
		if match != tt.out {
			t.Errorf("IsRegexpMatch(%s, %s) returned %v expected %v",
				tt.inS, tt.inP, match, tt.out)
		}
	}
}
