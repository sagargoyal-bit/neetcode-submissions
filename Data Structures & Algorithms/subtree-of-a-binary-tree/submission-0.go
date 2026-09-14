/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil{
		return false
	}
	if subRoot== nil{
		return true
	}
    // Check whether the trees match starting at this node.
    if sameTree(root, subRoot) {
        return true
    }  
	 // Otherwise, search the left and right subtrees.
    return isSubtree(root.Left, subRoot) ||
        isSubtree(root.Right, subRoot)

}

func sameTree(p *TreeNode, q *TreeNode) bool {
	//check both nodes
	if p ==nil && q ==nil{
		return true
	}
	 // Only one node is empty.
    if p == nil || q == nil {
        return false
    }
	//check the equal values
	if p.Val != q.Val{
		return false
	}

	left:=sameTree(p.Left,q.Left)
	right:=sameTree(p.Right,q.Right)

	return left && right

}