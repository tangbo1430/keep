package go支持的类型转换

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)
}
