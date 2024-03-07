package 数组_字符串

import (
	"fmt"
	"math/rand"
	"testing"
)

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 0),
		indices: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indices[val]
	if !ok {
		return false
	}
	this.nums[index] = this.nums[len(this.nums)-1] // 将要删除的元素放最后
	this.indices[this.nums[index]] = index         // 更新索引map，将最后一个元素的索引位置更新为要删的这个
	this.nums = this.nums[:len(this.nums)-1]       // 删除最后一个元素
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func TestRandom(t *testing.T) {
	c := Constructor()
	fmt.Println(c.Insert(0))
	fmt.Println(c.Insert(1))
	fmt.Println(c.Remove(0))
	fmt.Println(c.Insert(2))
	fmt.Println(c.Remove(1))
	fmt.Println(c.GetRandom())
}
