package datastruct

type NumArray []int

func Constructor(nums []int) NumArray {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

func (s NumArray) SumRange(left int, right int) int {
	return s[right+1] - s[left]
}
