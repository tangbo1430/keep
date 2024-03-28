package 二叉树

// 给定一个二叉树 root ，返回其最大深度。
//
//二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	res := max(left, right) + 1
	return res
}
