package 数组_字符串

// 你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度

func removeElement(nums []int, val int) int {
	left := 0
	for _, item := range nums {
		if item != val {
			nums[left] = item
			left++
		}
	}
	return left
}
