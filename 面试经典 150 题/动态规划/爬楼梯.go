package 动态规划

// 思路：1 2 3 5 8 想象是一个滚动的数组，滚一圈，第三个数是第一个数+第二个数
func climbStairs(n int) int {
	left, right, r := 0, 0, 1
	for i := 0; i < n; i++ {
		left = right
		right = r
		r = left + right
	}
	return r
}
