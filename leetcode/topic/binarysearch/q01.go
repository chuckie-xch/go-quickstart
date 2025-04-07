package binarysearch

func search(nums []int, target int) int {
	i := lowerBound(nums, target)
	if i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}

func lowerBound(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
