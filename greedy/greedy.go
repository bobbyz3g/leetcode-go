package greedy

import "strconv"

func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		var gasSum, costSum, count int
		// find out the next reachable index
		for count < n {
			j := (count + i) % n
			gasSum += gas[j]
			costSum += cost[j]
			// can not reach the j
			if costSum > gasSum {
				break
			}
			count++
		}
		// can complete circuit
		if count == n {
			return i
		} else {
			i += count + 1
		}
	}
	return -1
}

func MaximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	j, k, maxVal := -1, -1, len(s)-1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > s[maxVal] {
			maxVal = i
		} else if s[i] < s[maxVal] {
			j, k = i, maxVal
		}
	}
	if j == -1 {
		return num
	}
	s[j], s[k] = s[k], s[j]
	num, _ = strconv.Atoi(string(s))
	return num
}
