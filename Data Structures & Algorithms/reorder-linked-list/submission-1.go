/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	
    mid:=FindMid(head)
	secondPart:=mid.Next
	mid.Next=nil

	reversedlist:=reverse(secondPart)

	first:=head
	second:=reversedlist
	
    for second !=nil {
		firstNext :=first.Next
		secondNext :=second.Next

		first.Next=second
		second.Next=firstNext

		first=firstNext
		second=secondNext
	}

}

func FindMid(head *ListNode)*ListNode{
	fast:=head
	slow:=head
	for fast.Next !=nil && fast.Next.Next !=nil{
		fast =fast.Next.Next
		slow=slow.Next
	}
	return slow
}

func reverse(head *ListNode)*ListNode{
	var prev *ListNode
	current:=head
	for current !=nil{
		next:=current.Next
		current.Next=prev

		prev=current
		current=next
	}
	return prev
}