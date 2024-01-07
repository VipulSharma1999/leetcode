package main

func maxDepth(root *TreeNode) int {
	left, right := 0, 0
	if root == nil {
		return 0
	}
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}
	return 1 + maxx(left, right)
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
