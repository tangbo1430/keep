package 二叉树

// 深度优先遍历，自底向上变根节点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftNode := invertTree(root.Right)
	rightNode := invertTree(root.Left)

	root.Left = leftNode
	root.Right = rightNode
	return root
}
