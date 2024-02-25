package golang的struct能不能比较_

import (
	"fmt"
	"testing"
)

// Q:golang的struct能不能比较？

// 分情况
type T1 struct {
	a int
	b string
	c byte
	//d []int
}

type T2 struct {
	a int
	b string
	c byte
}

func Test(t *testing.T) {
	t1 := T1{
		a: 1,
		b: "1",
		c: 1,
	}
	t11 := T1{
		a: 1,
		b: "1",
		c: 1,
	}
	//t2 := T2{
	//	a: 0,
	//	b: "",
	//	c: 0,
	//}
	// 两个对象都是同一struct类型，且当struct中的类型都是可比类型时，就阔以比较
	fmt.Println(t1 == t11)

	//fmt.Println(t1 == t2)
}
