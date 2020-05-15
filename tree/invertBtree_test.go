package playgo

import (
	"fmt"
	"testing"
)

var one = &TreeNode{Val: 4,
	Left: &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	Right: &TreeNode{Val: 6,
		Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
		Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
}
var two = &TreeNode{Val: 1, Left: nil, Right: nil}
var three = &TreeNode{Val: 4,
	Left: &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	Right: &TreeNode{Val: 5, Left: nil, Right: nil}}

var four = &TreeNode{Val: 4,
	Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	Right: &TreeNode{Val: 5, Left: nil, Right: nil}}

var five = &TreeNode{Val: 4,
	Left: &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	Right: nil}

var six = &TreeNode{Val: 4,
	Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
	Right: nil}

var seven = &TreeNode{Val: 4,
	Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: nil},
		Right: nil},
	Right: &TreeNode{Val: 5,
		Left:  nil,
		Right: &TreeNode{Val: 9, Left: nil, Right: nil}}}

var oneP = &TreeNode{Val: 4,
	Right: &TreeNode{Val: 2,
		Right: &TreeNode{Val: 1, Left: nil, Right: nil},
		Left:  &TreeNode{Val: 3, Left: nil, Right: nil}},
	Left: &TreeNode{Val: 6,
		Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		Left:  &TreeNode{Val: 7, Left: nil, Right: nil}},
}
var twoP = &TreeNode{Val: 1, Left: nil, Right: nil}
var threeP = &TreeNode{Val: 4,
	Right: &TreeNode{Val: 2,
		Right: &TreeNode{Val: 1, Left: nil, Right: nil},
		Left:  &TreeNode{Val: 3, Left: nil, Right: nil}},
	Left: &TreeNode{Val: 5, Left: nil, Right: nil}}

func TestTreeHeight(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out int
	}{
		{one, 3},
		{two, 1},
		{three, 3},
		{four, 4},
		{five, 3},
		{six, 2},
		{seven, 4},
	}

	for _, tt := range tests {
		PrintTree(tt.in)
		if got := TreeHeight(tt.in); got != tt.out {
			t.Errorf("TreeHeight for (%v) want %d got %d",
				*tt.in, tt.out, got)
		}
	}
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out *TreeNode
	}{
		{one, oneP},
		{two, twoP},
		{three, threeP},
	}

	for _, tt := range tests {
		PrintTree(tt.in)
		if got := InvertTree(tt.in); CompareTree(got, tt.out) {
			fmt.Printf("Want:\n")
			PrintTree(tt.out)
			fmt.Printf("Got:\n")
			PrintTree(got)
			t.Errorf("InvertTree for (%#v) want %#v got %#v",
				tt.in, tt.out, got)
		}
	}

}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out bool
	}{
		{one, true},
		{two, true},
		{three, true},
		{nil, true},
		{four, false},
		{five, false},
		{six, true},
		{seven, false},
	}

	for _, tt := range tests {
		if got := IsBalanced(tt.in); got != tt.out {
			t.Errorf("IsBalanced for (%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}

}
