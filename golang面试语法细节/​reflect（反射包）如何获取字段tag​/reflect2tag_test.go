package _reflect_反射包_如何获取字段tag_

import (
	"fmt"
	"reflect"
	"testing"
)

// Q:reflect（反射包）如何获取字段tag？为什么json包不能导出私有变量的tag？
// json包里使用的时候，会在结构体里的字段边上加tag，有没有什么办法可以获取到这个tag的内容呢

// A：tag信息可以通过反射（reflect包）内的方法获取，通过一个例子加深理解
type J struct {
	a string //小写无tag
	b string `json:"B"` //小写+tag
	C string //大写无tag
	D string `json:"DD" otherTag:"good"` //大写+tag
}

func printTag(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%v个字段 %v 对应的json tag是 %v , 还有otherTag？ = %v \n", i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("otherTag"))
	}
}

func Test(t *testing.T) {
	j := J{
		a: "1",
		b: "2",
		C: "3",
		D: "4",
	}
	printTag(&j)
}
