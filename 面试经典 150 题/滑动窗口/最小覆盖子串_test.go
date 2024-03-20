package 滑动窗口

import "math"

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。

// 思路：滑动窗口，如果满足最小字串，就从左缩进窗口，如果不满足，就右张开窗口，
func minWindow(s string, t string) string {
	left, right := 0, 0
	start, length := 0, math.MaxInt64
	need := make(map[byte]int)
	window := make(map[byte]int)
	size := 0

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for right < len(s) {
		c := s[right]
		right++
		// 只记需要的字符
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				size++
			}
		}
		// 为什么取 len(need)长度，因为存在相同字符
		for size == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				// 如果此时这个字符满足了所需字符串长度，则要--
				if window[d] == need[d] {
					size--
					window[d]--
				} else {
					// 如果本来就没满足，也就不size--
					window[d]--
				}
			}
		}
	}

	if length != math.MaxInt64 {
		return s[start : start+length]
	}

	return ""
}
