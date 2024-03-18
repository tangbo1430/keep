package 双指针

// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
// 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

// 思路：双指针，前后夹击，如果两指针所指值加起来比target大，就缩左指针，如果小就缩右指针，因为条件给定该数组已按非递减顺序排列
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}
