/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	list1:=l1
	list2:=l2
	carry:=0
	for list1 !=nil && list2 !=nil{
		sum:=list1.Val+list2.Val+carry
		
		carry=sum/10
		digit:=sum%10

		current.Next=&ListNode{Val:digit}
		current=current.Next

		list1=list1.Next
		list2=list2.Next
	}
		for list1 !=nil{
		sum:=list1.Val+carry
		
		carry=sum/10
		digit:=sum%10
		
		current.Next=&ListNode{Val:digit}	
		current=current.Next
		
		list1=list1.Next
		}

		for list2 !=nil{
		sum:=list2.Val+carry
		
		carry=sum/10
		digit:=sum%10
		
		current.Next=&ListNode{Val:digit}	
		current=current.Next
		
		list2=list2.Next
		}

	
	if carry !=0{
		newNode:=&ListNode{Val:carry}
		current.Next=newNode
	}
	return dummy.Next

}
