package 链表

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: &ListNode{},
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Size {
		return -1
	}

	cur := this.Head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

// 头插法
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

// 尾插法
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	pre := this.Head
	this.Size++

	// 如果是-1说明在头节点插入
	index = max(index, 0)

	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	insertNode := &ListNode{
		Val:  val,
		Next: pre.Next,
	}

	pre.Next = insertNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	this.Size--
}
