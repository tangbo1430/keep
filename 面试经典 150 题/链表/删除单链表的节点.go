package 链表

// 删除指定节点：以前一个节点为基点，看后面一个是不是等于val，相等的话就删除pre后面那个节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	pre := head
	for pre.Next != nil && pre.Next.Val != val {
		pre = pre.Next
	}

	// 如果没找到val，此时pre已经走到链表尾部，不会触发此条件
	if pre.Next != nil {
		// 删除pre后面的节点
		pre.Next = pre.Next.Next
	}
	return head
}
