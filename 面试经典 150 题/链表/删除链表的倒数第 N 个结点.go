package 链表

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 思路，因为是找倒数第几个，所以指针p1先走n步，然后停，指针p2跟p1一起走，p1走到底了，p2就到了倒数第n个位置
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	removeNode := findRemoveNode(dummyNode, n+1)
	removeNode.Next = removeNode.Next.Next
	return dummyNode.Next
}

func findRemoveNode(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p2
}
