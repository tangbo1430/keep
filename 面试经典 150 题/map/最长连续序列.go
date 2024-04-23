package _map

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9

// 思路：用一个map存住所有数,然后遍历找item-1不在map的数字，然后二次遍历看它+1在不在map，在就继续往下找
func longestConsecutive(nums []int) int {
	longestStreak := 0

	seen := make(map[int]bool)
	for _, item := range nums {
		seen[item] = true
	}

	for item := range seen {
		// 如果item-1不在map，就尝试找item+1有没有，有就说明是连续的
		if !seen[item-1] {
			currentNum := item
			currentStreak := 1
			// 只有当一个数是连续序列的第一个数的情况下才会进入内层循环
			for seen[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
