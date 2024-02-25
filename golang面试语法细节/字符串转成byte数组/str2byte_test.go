package 字符串转成byte数组

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// Q：字符串转成byte数组，会发生内存拷贝吗？

// 字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝。那么问题来了。

// 频繁的内存拷贝操作听起来对性能不大友好。有没有什么办法可以在字符串转成切片的时候不用发生拷贝呢？

func Test(t *testing.T) {
	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b)
}
