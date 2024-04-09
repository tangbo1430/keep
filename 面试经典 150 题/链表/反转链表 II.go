package 链表

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ret := &ListNode{Val: -1}
	ret.Next = head
	pre := ret

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		nextNode := cur.Next
		cur.Next = nextNode.Next
		nextNode.Next = pre.Next
		pre.Next = nextNode
	}

	return ret.Next
}
