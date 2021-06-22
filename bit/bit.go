package bit

// HammingWeight returns the hamming weight of a uint32 number.
func HammingWeight(num uint32) int {
	var nums uint32
	for num > 0 {
		nums += (num & 1)
		num = num >> 1
	}
	return int(nums)
}

// ReverseBits revers all bits of given number.
func ReverseBits(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res |= num & 1 << (31 - i)
		num >>= 1
	}

	return res
}

// CountBits returns an array ans of length n + 1 such that for each i (0 <= i <= n),
// res[i] is the number of 1's in the binary representation of i.
func CountBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}
