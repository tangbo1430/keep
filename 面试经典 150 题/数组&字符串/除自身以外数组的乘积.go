package 数组_字符串

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]

// left:  1 1 2 6
// right: 1 4 12 24

// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]

// 思路：因为要O(n)，所以可以提前算好index左边和右边的乘积
// 两次遍历，可以去除头和尾之后的index左侧和右侧乘积，然后再相乘，就是除index外的乘积
func productExceptSelf(nums []int) []int {
	left, right, result := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1 // 起始因为是要除自身外，然后×左边的数
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	right[len(nums)-1] = 1 // 起始因为是要除自身外，然后×右边的数
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for index, _ := range nums {
		result[index] = left[index] * right[index]
	}
	return result
}
