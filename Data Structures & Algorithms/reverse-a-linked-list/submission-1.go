/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

    var prev *ListNode = nil
	currentNode:= head
	for currentNode !=nil{
	 next:=currentNode.Next //save the next node
	 currentNode.Next=prev //reverse
	 prev=currentNode //move prev
	 currentNode=next //move next
	}
	return prev
}
