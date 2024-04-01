package 二叉树

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
