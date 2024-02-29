package 数组_字符串

// 投票方法，前n个数，非众数与众数会把votes抵消掉，如果votes没抵消到0，就说明那个数是众数
func majorityElement(nums []int) int {
	//quickSort(nums, 0, len(nums) - 1)
	//return nums[len(nums)/2]
	votes := 0
	moreNum := 0
	for _, item := range nums {
		if votes == 0 {
			moreNum = item
		}
		if item == moreNum {
			votes++
		} else {
			votes--
		}
	}
	return moreNum
}

// 也可以使用快排，排完了如何在找n/2一定是众数
func quickSort(nums []int, low, high int) {
	if low < high {
		i, j, tmp := low, high, nums[(low+high)/2]
		for i <= j {
			for nums[i] < tmp {
				i++
			}
			for nums[j] > tmp {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if low < j {
			quickSort(nums, low, j)
		}
		if high > i {
			quickSort(nums, i, high)
		}
	}
}
