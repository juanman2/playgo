package playgo

import (
	"fmt"
)

// TreeNode is the definition for a binary tree type
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeHeight returns the maximum height of the btree
func TreeHeight(r *TreeNode) int {
	if r == nil {
		return 0
	}

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

// IsBalanced returns true if the difference between the left and right
// subtrees of any node in the tree is less than one.
func IsBalanced(root *TreeNode) bool {
	bal, _ := isBalancedTree(root)
	return bal
}

// IsBalanced returns true if the difference between the left and right
// subtrees of any node in the tree is less than one. It returns
// the height of the tree to aid computation
func isBalancedTree(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	rb, rh := isBalancedTree(root.Right)
	lb, lh := isBalancedTree(root.Left)
	if rb == false || lb == false {
		return false, 0
	}

	delta := lh - rh
	if delta < -1 || delta > 1 {
		return false, 0
	}

	height := lh
	if rh > height {
		height = rh
	}

	return true, height + 1
}

func printTreeLine(n *TreeNode, l, level int, line *[]string) {

	if l == level {
		if n == nil {
			*line = append(*line, "- ")
		} else {
			*line = append(*line, fmt.Sprintf("%d ", n.Val))
		}
	}

	if l < level {
		if n == nil {
			*line = append(*line, "- ")
			*line = append(*line, "- ")
		} else {
			printTreeLine(n.Left, l+1, level, line)
			printTreeLine(n.Right, l+1, level, line)
		}
	}
}

// PrintTree is a pretty print for a tree
func PrintTree(n *TreeNode) {

	h := TreeHeight(n)
	if h > 4 {
		fmt.Printf("Warning: this function will print at most 4 levels")
	}

	for l := 0; l < h && l < 4; l = l + 1 {
		line := make([]string, 0)
		printTreeLine(n, 0, l, &line)
		switch l {
		case 0:
			fmt.Printf("       %s\n", line[0])
		case 1:
			fmt.Printf("      / \\  \n")
			fmt.Printf("     %s  %s \n", line[0], line[1])
		case 2:
			fmt.Printf("   / \\   / \\\n")
			fmt.Printf("  %s  %s %s %s\n",
				line[0], line[1], line[2], line[3])
		case 3:
			fmt.Printf(" /\\  /\\  /\\  /\\\n")
			fmt.Printf("%s%s%s%s %s%s%s%s\n",
				line[0], line[1], line[2], line[3],
				line[4], line[5], line[6], line[7])
		}
	}
	fmt.Println()

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
