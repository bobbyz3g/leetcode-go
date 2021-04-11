package numbers

var factors = []int{2, 3, 5}

// IsUgly returns true if n is a ugly number.
// Ugly number is a positive number whose prime factors only include 2, 3, and/or 5
func IsUgly(n int) bool {
	if n < 0 {
		return false
	}

	for _, v := range factors {
		for n%v == 0 {
			n /= v
		}
	}

	return n == 1
}

// NthUglyNumber returns the nth ugly number.
// Ugly number is a positive number whose prime factors only include 2, 3, and/or 5
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
