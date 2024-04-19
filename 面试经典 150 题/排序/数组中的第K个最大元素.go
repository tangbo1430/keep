package 排序

// 快排
func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, left, right int) {
	if left < right {
		i, j, tmp := left, right, nums[(right+left)/2]
		for i <= j {
			for tmp > nums[i] {
				i++
			}
			for tmp < nums[j] {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if j > left {
			quickSort(nums, left, j)
		}
		if right > i {
			quickSort(nums, i, right)
		}
	}
}
