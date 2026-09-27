/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy :=&ListNode{}
	currnetList:=dummy
    for list1 !=nil && list2 !=nil{
		if list1.Val <= list2.Val{
			currnetList.Next=list1
			list1=list1.Next
		}else{
			currnetList.Next=list2	
			list2=list2.Next	
		}
		currnetList=currnetList.Next	
	}
	if list1 !=nil{
		currnetList.Next =list1
	}
	if list2 !=nil{
		currnetList.Next =list2
	}
	return dummy.Next
}
