package json包变量不加tag会怎么样

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 转为json后首字母小写的不管加不加tag都不能转为json里的内容，
// 而大写的加了tag可以取别名，不加tag则json内的字段跟结构体字段原名一致。

type J struct {
	a string //小写无tag
	b string `json:"B"` //小写+tag
	C string //大写无tag
	D string `json:"DD"` //大写+tag
}

func Test(t *testing.T) {
	j := J{
		a: "1",
		b: "2",
		C: "3",
		D: "4",
	}
	fmt.Printf("转为json前j结构体的内容 = %+v\n", j)
	jsonInfo, _ := json.Marshal(j)
	fmt.Printf("转为json后的内容 = %+v\n", string(jsonInfo))
}
