package 二分查找

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 示例 1:
//输入: nums = [1,3,5,6], target = 5
//输出: 2

//示例 2:
//输入: nums = [1,3,5,6], target = 2
//输出: 1

// 思路：二分查找，默认位置给数组长度，然后按照target二分查
func searchInsert(nums []int, target int) int {
	ans := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
