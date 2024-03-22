package _map

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

// 思路：用map存已经出现过的数字与对应index，如果target-当前数的结果在map中存在，则说明找到结果
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for i, item := range nums {
		if _, ok := memo[target-item]; ok {
			return []int{memo[target-item], i}
		}
		memo[nums[i]] = i
	}
	return []int{}
}
