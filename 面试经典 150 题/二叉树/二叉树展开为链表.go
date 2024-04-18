package 二叉树

// 思路：前序遍历把节点存到切片，然后遍历切片，做成链表，左子树为空就行
func flatten(root *TreeNode) {
	list := myPreorderTraversal(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left = nil
		pre.Right = cur
	}
	return
}

func myPreorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		// 前序遍历，左没有了再放右
		list = append(list, root)
		list = append(list, myPreorderTraversal(root.Left)...)
		list = append(list, myPreorderTraversal(root.Right)...)
	}
	return list
}
