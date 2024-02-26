package 能说说uintptr和unsafe_Pointer的区别吗

import (
	"fmt"
	"testing"
	"unsafe"
)

// Q:能说说uintptr和unsafe.Pointer的区别吗

// unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
// unsafe.Pointer 可以和 普通指针 进行相互转换；
// unsafe.Pointer 可以和 uintptr 进行相互转换

type W struct {
	b int32
	c int64
}

func Test(t *testing.T) {
	var w *W = new(W)
	// 这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	// uintptr(unsafe.Pointer(w)) 获取了 w 的指针起始值
	// unsafe.Offsetof(w.b) 获取 b 变量的偏移量
	// 两个相加就得到了 b 的地址值，将通用指针 Pointer 转换成具体指针 ((*int)(b))，
	// 通过 * 符号取值，然后赋值。*((*int)(b)) 相当于把 (*int)(b) 转换成 int 了，最后对变量重新赋值成 10，这样指针运算就完成了。

	// 现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	//*((*int)(b)) = 10
	*(*int)(b) = 10
	// 此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}
