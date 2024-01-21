package greedy

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
