package bits

func HammingWeight(num uint32) int {
	var nums uint32
	for num > 0 {
		nums += (num & 1)
		num = num >> 1
	}
	return int(nums)
}
