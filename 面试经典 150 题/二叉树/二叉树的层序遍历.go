package 二叉树

var ans [][]int

// bfs和dfs都能做
func levelOrder(root *TreeNode) [][]int {
	ans = make([][]int, 0)
	dfs(root, 0)
	return ans
}

// 深度优先
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

// 广度优先
func bfs(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	outTmp := []*TreeNode{root}
	for level := 0; len(outTmp) > 0; level++ {
		inTmp := []*TreeNode{}
		ans = append(ans, []int{})
		for j := 0; j < len(outTmp); j++ {
			node := outTmp[j]
			ans[level] = append(ans[level], node.Val)
			if node.Left != nil {
				inTmp = append(inTmp, node.Left)
			}
			if node.Right != nil {
				inTmp = append(inTmp, node.Right)
			}
		}
		outTmp = inTmp
	}
	return ans
}
