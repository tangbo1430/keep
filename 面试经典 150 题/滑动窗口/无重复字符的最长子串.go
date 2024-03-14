package 滑动窗口

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

// 脑补一个滑动的窗口，遇到重复的，左边往右边缩一个，不重复的，右边往左边扩进来
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	res := 0
	memo := make(map[byte]int, 0)

	for r < len(s) {
		c := s[r]
		r++
		memo[c]++
		// 遇到重复的了
		for memo[c] > 1 {
			d := s[l]
			l++
			memo[d]--
		}
		res = max(res, r-l)
	}
	return res
}
