package 二叉树

// 前序遍历，左中右
func preorderTraversal(root *TreeNode) (ans []int) {
	var preFunc func(node *TreeNode)
	preFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preFunc(node.Left)
		preFunc(node.Right)
	}

	preFunc(root)
	return
}
