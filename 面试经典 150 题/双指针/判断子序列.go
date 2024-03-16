package 双指针

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

// 输入：s = "abc", t = "ahbgdc"
// 输出：true

// 输入：s = "axc", t = "ahbgdc"
// 输出：false

// 思路：双指针，左右指针，如果遇到相同的，指针就推进，到最后左指针长度跟s字符串长度一样，说明是字串
func isSubsequence(s string, t string) bool {
	l, r := 0, 0
	m, n := len(s), len(t)
	for l < m && r < n {
		if s[l] == t[r] {
			l++
		}
		r++
	}
	return m == l
}
