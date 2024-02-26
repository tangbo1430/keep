package JSON_标准库对_nil_slice_和_空_slice_的处理是_致的吗__

import (
	"fmt"
	"testing"
)

// ⾸先JSON 标准库对 nil slice 和 空 slice 的处理是不⼀致.
// 通常错误的⽤法，会报数组越界的错误，因为只是声明了slice，却没有给实例化的对象。

func Test(t *testing.T) {
	// 会导致 index out of range [1] with length 0
	//var slice []int
	//slice[1] = 0

	slice := make([]int, 0)
	//slice := []int{}
	fmt.Println(slice)
}
