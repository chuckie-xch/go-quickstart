package binarysearch

func maximumCount(nums []int) int {
	n := len(nums)
	lowerIndex := lowerIndex(nums, n)
	higherIndex := higherIndex(nums, n)
	return max(n-lowerIndex, higherIndex+1)
}

func lowerIndex(nums []int, n int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func higherIndex(nums []int, n int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
