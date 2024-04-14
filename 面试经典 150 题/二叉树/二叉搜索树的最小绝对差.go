package 二叉树

import "math"

// 思路：二叉搜索树的中序遍历是升序的，按照中序遍历，前一个节点val-后一个节点val找出最小值即可
func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1

	var dfs func(r *TreeNode)

	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r.Left)

		if pre != -1 && r.Val-pre < ans {
			ans = r.Val - pre
		}
		pre = r.Val
		dfs(r.Right)
	}

	dfs(root)
	return ans
}
