package 数组_字符串

import (
	"fmt"
	"testing"
)

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置

// 输入: nums = [2,3,0,1,4]
// 输出: 2
func jump(nums []int) int {
	maxJump := 0 // 最大跳跃距离
	board := 0   // 跳跃边界，到达边界时只能跳
	step := 0    // 跳跃次数
	for i := 0; i < len(nums)-1; i++ {
		// 取最大跳跃距离
		maxJump = max(maxJump, nums[i]+i)
		// 如果到达边界就起跳
		if i == board {
			// 更新边界下标
			board = maxJump
			step++
		}
		fmt.Println("board:", board)
		fmt.Println("maxJump:", maxJump)
	}
	return step
}

func TestJump(t *testing.T) {
	nums := []int{1, 2, 1, 1, 4}
	jump(nums)
}
