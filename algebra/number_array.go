package algebra

// NumArray is used to calculate range sum of array.
type NumArray struct {
	preSum []int
}

// Constructor returns a new NumArray.
func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)

	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}

	return NumArray{
		preSum,
	}
}

// SumRange returns then sum of range (1, j).
func (n *NumArray) SumRange(i int, j int) int {
	return n.preSum[j+1] - n.preSum[i]
}
