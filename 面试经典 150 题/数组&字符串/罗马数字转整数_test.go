package 数组_字符串

import (
	"fmt"
	"testing"
)

// 输入: s = "III"
// 输出: 3

// 输入: s = "LVIII"
// 输出: 58
// 解释: L = 50, V= 5, III = 3.

// 输入: s = "IV"
// 输出: 4

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			// 如果是一直前一个小于后一个，那么最终结果就是最大的减去前面的
			ans += value
		}
	}
	return
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("IVX"))
}
