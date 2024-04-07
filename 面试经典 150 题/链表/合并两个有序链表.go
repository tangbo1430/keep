package 链表

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	nextNode := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			nextNode.Next = list1
			list1 = list1.Next
		} else {
			nextNode.Next = list2
			list2 = list2.Next
		}
		nextNode = nextNode.Next
	}
	if list1 != nil {
		nextNode.Next = list1
	}
	if list2 != nil {
		nextNode.Next = list2
	}
	return head.Next
}
