package 数组_字符串

import (
	"testing"
)

// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
// 你需要按照以下要求，给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
// 输入：ratings = [1,0,2]
// 输出：5
// 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。

// 思路：需要满足两种匹配规则：
// 1.左匹配，如果某孩子小于左边的孩子评分，则孩子+1，遍历一遍得出，左边匹配规则
// 2.右匹配，如果某小孩大于右边的孩子评分，则孩子+1，遍历一遍得出，右边匹配规则
// 3.两个匹配规则得到的最大值，累计就是最小需要发放的糖果
// 为什么要遍历获得左右匹配规则，因为单方向，只知道相邻左侧或右侧是不是比自己大，不知道相隔的孩子评分

func candy(ratings []int) int {
	res := 0
	left := make([]int, len(ratings))
	// 第i号孩子比他左边的孩子评分高，则第i号孩子比第i-1号孩子多一个糖果
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		// 第i号孩子比他右边的孩子评分高，则第i号孩子比第i+1号孩子多一个糖果
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		// 第i号孩子最少需要max(right,left[i])个糖果才能满足条件
		res += max(right, left[i])
	}
	return res
}

func TestCandy(t *testing.T) {
	candy([]int{1, 2, 2})
}
