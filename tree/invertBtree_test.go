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
