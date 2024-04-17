package 回溯

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 思路：回溯

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, visited []int, res *[][]int) {
	// 当达到排列需要的长度
	if len(nums) == len(visited) {
		tmp := make([]int, len(nums))
		copy(tmp, visited)
		*res = append(*res, tmp)
		// 结束本次排列
		return
	}

	for i := 0; i < len(nums); i++ {
		var seen bool
		for _, item := range visited {
			if item == nums[i] {
				seen = true
				break
			}
		}
		if seen {
			continue
		}

		visited = append(visited, nums[i])
		dfs(nums, visited, res)
		visited = append(visited, nums[len(nums)-1])
	}
}
