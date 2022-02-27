package str

// CheckInclusion returns true if s2 contains the permutation of s1.
func CheckInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	var target, cur [26]int
	for _, v := range s1 {
		target[v-'a']++
	}

	for i := 0; i < l1; i++ {
		cur[s2[i]-'a']++
	}

	if cur == target {
		return true
	}

	for i := l1; i < l2; i++ {
		cur[s2[i-l1]-'a']--
		cur[s2[i]-'a']++
		if cur == target {
			return true
		}
	}

	return false
}

// LongestCommonPrefix returns the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {
	var result string
	if len(strs) == 0 {
		return result
	}

	result = strs[0]
	for i := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return result
}

// LongestPalindrome return the longest palindromic substring in s
func LongestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	begin, maxLen := 0, 1

	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			if dp[i][j] && j-i+1 > maxLen {
				begin = i
				maxLen = j - i + 1
			}
		}
	}

	return s[begin : begin+maxLen]
}

// Partition returns all possible palindrome partitioning of s.
func Partition(s string) [][]string {
	l := len(s)

	// flags of str[i, j] is palindrome string or not.
	// f(i, j) = true when  i >= j
	// f(i, j) = f(i+1, j-1) && str[i] == str[j] when i < j
	pArr := make([][]bool, l)

	// initial pArr
	for i := range pArr {
		pArr[i] = make([]bool, l)
		for j := range pArr[i] {
			pArr[i][j] = true
		}
	}

	// check f(i, j)
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			pArr[i][j] = s[i] == s[j] && pArr[i+1][j-1]
		}
	}

	result := make([][]string, 0, 1)
	splits := []string{}

	// dfs will be used in closures, you need to declare first
	var dfs func(int)

	dfs = func(i int) {
		// no more elements, ends recursion
		if i == l {
			result = append(result, append([]string(nil), splits...))
			return
		}
		for j := i; j < l; j++ {
			if pArr[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)

	return result
}

// LengthOfLongestSubstring returns the length of the longest
// substring without repeating characters.
// link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
