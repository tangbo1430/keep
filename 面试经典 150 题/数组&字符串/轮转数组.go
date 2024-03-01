package 数组_字符串

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]

// 思路：翻转三次即可，先全部翻转，再翻转 前k%len(num)个，再翻转k%len(num) ~ len(num)个
func rotate(nums []int, k int) {
	// 取模，因为入参可能会超过数组本身长度
	index := k % len(nums)
	Reverse(nums)
	Reverse(nums[:index])
	Reverse(nums[index:])
	return
}

func Reverse(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums
}
