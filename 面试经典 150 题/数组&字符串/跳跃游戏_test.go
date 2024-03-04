package 数组_字符串

import (
	"fmt"
	"testing"
)

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
func canJump(nums []int) bool {
	jumpMax := 0
	for i := 0; i < len(nums); i++ {
		// 当前下标大于最大跳跃距离
		if i > jumpMax {
			return false
		}
		// 遍历刷新最大能跳多远，最远=当前下标+当前或之前最大跳跃距离
		jumpMax = max(jumpMax, i+nums[i])
		fmt.Println(jumpMax)
	}
	return true
}

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	canJump(nums)
}
