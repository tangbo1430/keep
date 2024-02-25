package 翻转含有中文_数字_英文字母的字符串

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	src := "你好abc啊哈哈"
	dst := reverse([]rune(src))
	fmt.Printf("%v\n", string(dst))
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
