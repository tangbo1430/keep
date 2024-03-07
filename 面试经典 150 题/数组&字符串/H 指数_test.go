package 数组_字符串

import (
	"fmt"
	"testing"
)

// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
//
//                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。
// 如果 h 有多种可能的值，h 指数 是其中最大的那个。

// 输入：citations = [3,0,6,1,5]
// 输出：3
// 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
// 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。

// 思路：不要被题目混淆了，本质就是再找数组下标>=（大于等于数组下标的元素），如果满足就返回

// 计数
func hIndex(citations []int) int {
	n := len(citations)
	h := 0
	count := make([]int, n+1)
	for _, item := range citations {
		count[min(n, item)]++ // min(n, item)记录每个元素出现的次数，放置在count数组元素值对应下标上
		fmt.Println(count)
	}

	for i := n; ; i-- {
		h += count[i] // 引用次数记录
		if h >= i {   // 如果引用次数大于当前下标，则达到目标
			return i
		}
	}
}

// 本质是找某个下标上的数，大于等于整个数组的元素的数量，是不是大于下标
// 二分搜索
func hIndex_(citations []int) int {
	// 答案最多只能到数组长度
	left, right := 0, len(citations)
	var mid int
	for left < right {
		// +1 防止死循环
		mid = (left + right + 1) / 2
		fmt.Println("mid:", mid)
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		fmt.Println("cnt:", cnt)
		if cnt >= mid { // 如果当前元素，满足当前下标大于等于（下标元素大于等于全部元素中的数量），说明后续还有可能有更大的
			// 要找的答案在 [mid,right] 区间内
			left = mid
		} else { // 如果当前元素，满足当前下标小于（下标元素小于全部元素中的数量），说明后续还有可能有更小的
			// 要找的答案在 [0,mid) 区间内
			right = mid - 1
		}
		fmt.Println(left, right)

	}
	return left
}

func TestHIndex(t *testing.T) {
	data := []int{1, 4, 4, 4, 5}
	fmt.Println(hIndex_(data))
}
