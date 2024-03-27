package 二分查找

import "math"

// 峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

// 示例 1：
//
//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。

// 示例 2：
//
//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。

// 思路：二分查找，写一个内部fun取数，然后比两边大小就行
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	f := func(n int) int {
		if n == -1 || n == len(nums) {
			return math.MinInt64
		}
		return nums[n]
	}

	for {
		mid := (right + left) / 2
		if f(mid) > f(mid-1) && f(mid) > f(mid+1) {
			return mid
		}
		if f(mid) < f(mid-1) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}
