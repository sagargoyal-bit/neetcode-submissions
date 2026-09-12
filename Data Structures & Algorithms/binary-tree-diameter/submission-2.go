/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	diameterVal := 0
	var depth func(root *TreeNode) int

	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left)
		right := depth(root.Right)
		diameterVal = max(diameterVal, left+right)
		//calculate the height
		return 1 + max(left, right)
	}
	depth(root)
	return diameterVal
}
