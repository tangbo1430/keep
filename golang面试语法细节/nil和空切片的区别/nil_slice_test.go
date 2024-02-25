package nil和空切片的区别

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// Q： nil和空切片的区别

// 1.nil切片和空切片指向的地址不一样。nil空切片引用数组指针地址为0（无指向任何实际地址）
// 2.空切片的引用数组指针地址是有的，且固定为一个值

func Test(t *testing.T) {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
