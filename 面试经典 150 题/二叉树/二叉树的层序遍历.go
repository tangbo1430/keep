package 二叉树

var ans [][]int

// bfs和dfs都能做
func levelOrder(root *TreeNode) [][]int {
	ans = make([][]int, 0)
	dfs(root, 0)
	return ans
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	// 到达下一层初始化二维数组
	if len(ans) == level {
		ans = append(ans, []int{})
	}

	ans[level] = append(ans[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func bfs() {

}
