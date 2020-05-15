package playgo

import (
	"fmt"
	"math"
	"strings"
)

// TreeNode is the definition for a binary tree type
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeHeight returns the maximum height of the btree
func TreeHeight(r *TreeNode) int {
	hl, hr := 0, 0
	if r.Left != nil {
		hl = TreeHeight(r.Left)
	}

	if r.Right != nil {
		hr = TreeHeight(r.Right)
	}

	if hl > hr {
		return hl + 1
	}
	return hr + 1
}

func printTreeLine(n *TreeNode, l, level int, line *[]string) {

	if l == level {
		*line = append(*line, fmt.Sprintf("%d ", n.Val))
	}

	if l < level {
		if n.Left != nil {
			printTreeLine(n.Left, l+1, level, line)
		} else {
			*line = append(*line, "- ")
		}

		if n.Right != nil {
			printTreeLine(n.Right, l+1, level, line)
		} else {
			*line = append(*line, "- ")
		}
	}
}

// PrintTree is a pretty print for a tree
func PrintTree(n *TreeNode) {

	h := TreeHeight(n)
	if h > 5 {
		fmt.Printf("Warning: this function will print at most 6 levels")
	}

	for l := 0; l < h && l < 6; l = l + 1 {
		line := make([]string, 2>>l)
		printTreeLine(n, 0, l, &line)
		width := int(math.Exp2(float64(h)))
		if l%2 == 0 {
			width = width - 1
		}
		fmt.Printf("%*s\n", width, strings.Join(line, ""))
	}

}

func invertNode(n *TreeNode) {
	t := n.Left
	n.Left = n.Right
	n.Right = t
}

// inverTree solve the LeetCode problem
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	invertNode(root)
	return root
}

// InvertTree swap the left and right children on all the nodes
func InvertTree(root *TreeNode) *TreeNode {
	invertTree(root)
	PrintTree(root)
	return root
}

// CompareTree will compare all the nodes and make sure the values are same
func CompareTree(n *TreeNode, p *TreeNode) bool {

	if n.Val != p.Val {
		return false
	}

	leftE, rightE := false, false

	if n.Left != nil {
		if p.Left != nil {
			leftE = CompareTree(n.Left, p.Left)
		}
	}
	if leftE == false {
		return false
	}

	if n.Right != nil {
		if p.Right != nil {
			rightE = CompareTree(n.Right, p.Right)
		}
	}
	if rightE == false {
		return false
	}

	return true
}
