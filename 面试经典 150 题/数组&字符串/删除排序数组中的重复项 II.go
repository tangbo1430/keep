package 数组_字符串

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
