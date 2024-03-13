package 数组_字符串

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

// 思路：将所有字符串依次跟strs[0]比较，有相同的就index++
func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	i := 0
	for i < length && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}
