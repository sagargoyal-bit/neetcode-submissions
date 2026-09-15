/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    return checkBothTree(p,q)
}

func checkBothTree(p *TreeNode, q *TreeNode) bool {
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

	left:=checkBothTree(p.Left,q.Left)
	right:=checkBothTree(p.Right,q.Right)

	return left && right

}
