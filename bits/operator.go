package bits

func HammingWeight(num uint32) int {
	var nums uint32
	for num > 0 {
		nums += (num & 1)
		num = num >> 1
	}
	return int(nums)
}

func ReverseBits(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res |= num & 1 << (31 - i)
		num >>= 1
	}

	return res
}
