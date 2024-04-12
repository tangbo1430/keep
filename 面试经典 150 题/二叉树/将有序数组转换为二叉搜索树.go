package 二叉树

// 二叉搜索树的中序遍历是升序序列，题目给定的数组是按照升序排序的有序数组，因此可以确保数组是二叉搜索树的中序遍历序列。

// 方法：中序遍历，总是选择中间位置左边的数字作为根节点
func sortedArrayToBST(nums []int) *TreeNode {
	return ZSort(nums, 0, len(nums)-1)
}

// 中序遍历
func ZSort(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	root := &TreeNode{Val: nums[mid]}
	root.Right = ZSort(nums, mid+1, right)
	root.Left = ZSort(nums, left, mid-1)
	return root
}
