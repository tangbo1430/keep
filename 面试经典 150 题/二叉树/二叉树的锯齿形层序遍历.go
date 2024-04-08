package 二叉树

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
//（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	index := 0
	res := make([][]int, 0)
	out := []*TreeNode{root}
	for len(out) > 0 {
		tmp := make([]int, 0)
		res = append(res, []int{})
		in := []*TreeNode{}
		for i := 0; i < len(out); i++ {
			node := out[i]
			if index%2 == 0 {
				res[index] = append(res[index], node.Val)
			} else {
				tmp = append(tmp, node.Val)
			}

			if node.Left != nil {
				in = append(in, node.Left)
			}

			if node.Right != nil {
				in = append(in, node.Right)
			}
		}
		for len(tmp) > 0 {
			res[index] = append(res[index], tmp[len(tmp)-1])
			tmp = tmp[:len(tmp)-1]
		}

		out = in
		index++
	}
	return res
}
