package 数组_字符串

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	s := 1
	for f := 1; f < len(nums); f++ {
		if nums[f] != nums[f-1] {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}
