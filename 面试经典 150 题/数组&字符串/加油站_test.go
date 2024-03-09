package 数组_字符串

import (
	"testing"
)

// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

// 输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// 输出: 3
// 解释:
// 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
// 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
// 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
// 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
// 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
// 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
// 因此，3 可为起始索引。

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		// 记录当前位置出发的积累邮箱和花费，以及走了几步
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			// 模拟从当前位置出发，走cnt步，如果油箱空了，就结束循环
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			// 向前走一步
			cnt++
		}
		// 如果已经走了一圈，则返回出发原点
		if cnt == n {
			return i
		} else {
			// 如果没走到一圈，本来是index往下+1，但是有可能环绕了一圈，所以还要加上cnt
			i += cnt + 1
		}
	}
	return -1
}

func TestCanCompleteCircuit(t *testing.T) {
	canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
}
