package algebra

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)

	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}

	return NumArray{
		preSum,
	}
}

func (n *NumArray) SumRange(i int, j int) int {
	return n.preSum[j+1] - n.preSum[i]
}
