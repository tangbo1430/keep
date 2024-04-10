package 链表

// 思路：提前拿到后两个节点，然后判断值是否相等就行
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ret := &ListNode{0, head}

	cur := ret
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val != x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return ret.Next
}
