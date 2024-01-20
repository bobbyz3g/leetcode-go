package str

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

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

// LengthOfLastWord returns the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
func LengthOfLastWord(s string) int {
	var end, i int
	var found bool
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !found {
			end = i
			found = true
		}
		if s[i] == ' ' && found {
			break
		}
	}
	if found {
		return end - i
	}
	return 0
}

// AddBinary return sum of given string as a binary string.
func AddBinary(a string, b string) string {
	la, lb := len(a), len(b)
	if la > lb {
		return addBinary(a, b)
	} else {
		return addBinary(b, a)
	}
}

func addBinary(l string, s string) string {
	sumStr := make([]byte, len(l)+1)
	sumStr[0] = 48
	for i := 1; i < len(l)+1; i++ {
		sumStr[i] = l[i-1]
	}

	diff := len(sumStr) - len(s)
	// '0': 48, '1': 49, '2': 50
	var n uint8
	for i := len(s) - 1; i >= 0; i-- {
		sum := (s[i] - 48) + (sumStr[i+diff] - 48) + n
		n = sum / 2
		sumStr[i+diff] = sum%2 + 48
	}

	for i := diff - 1; i >= 0; i-- {
		sum := (sumStr[i] - 48) + n
		n = sum / 2
		sumStr[i] = sum%2 + 48
	}

	if sumStr[0] == '0' {
		return string(sumStr[1:])
	}
	return string(sumStr)
}

// IsPalindrome returns true if s is palindrome.
//
// It will convert all uppercase letters into lowercase
// letters and removing all non-alphanumeric characters.
func IsPalindrome(s string) bool {
	l := len(s)
	if l == 0 { // emtpy string is palindrome
		return true
	}
	i, j := 0, l-1
	for j > i {
		if !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
			continue
		}

		if !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j--
			continue
		}
		si, sj := s[i], s[j]
		if 'a' <= si && si <= 'z' {
			si = si - 32
		}
		if 'a' <= sj && sj <= 'z' {
			sj = sj - 32
		}
		if si != sj {
			return false
		}
		i++
		j--
	}
	return true
}

// ConvertToTitle returns corresponding column title as
// it appears in an Excel sheet.
// e.g.
// A -> 1, B -> 2 , C -> 3,Z -> 26
// AA -> 27, AB -> 28
func ConvertToTitle(column int) string {
	var res []rune
	for column > 0 {
		column--
		res = append(res, rune('A'+column%26))
		column /= 26
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// TitleToNumber returns its corresponding column number.
func TitleToNumber(title string) int {
	num := 0
	l := len(title)
	for i := 0; i < l; i++ {
		n := int(title[i] - 'A' + 1)
		p := math.Pow(26, float64(l-i-1))
		num += n * int(p)
	}
	return num
}

// IsIsomorphic returns true if all occurrences of a character in s can
// be replaced with another character while preserving the order of characters in t.
func IsIsomorphic(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// CheckIfPangram reports whether sentence is pangram.
func CheckIfPangram(sentence string) bool {
	state := 0
	for _, c := range sentence {
		state |= 1 << (c - 'a')
	}
	return state == 1<<26-1
}

// NumbersAscending returns true if all the numbers in
// s are strictly increasing from left to right.
// A sentence is a list of tokens separated by a single space with
// no leading or trailing spaces.
func NumbersAscending(s string) bool {
	tokens := strings.Split(s, " ")
	prev := -1
	for _, v := range tokens {
		cur, err := strconv.Atoi(v)
		if err != nil { // not a number
			continue
		}
		if cur <= prev {
			return false
		}
		prev = cur
	}
	return true
}

// DigitCount returns true if for every index i in the range 0 <= i < n,
// the digit i occurs num[i] times in num, otherwise return false.
func DigitCount(num string) bool {
	counter := make(map[int32]int32)
	for _, v := range num {
		n := v - '0'
		counter[n] = counter[n] + 1
	}
	for i, v := range num {
		n := v - '0'
		if counter[int32(i)] != n {
			return false
		}
	}
	return true
}

// Evaluate evaluates all the bracket pairs with key in knowledge.
func Evaluate(s string, knowledge [][]string) string {
	knowledgeMap := make(map[string]string, len(knowledge))
	for _, v := range knowledge {
		knowledgeMap[v[0]] = v[1]
	}
	builder := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] != '(' {
			builder.WriteByte(s[i])
			continue
		}
		// get index of ')'
		j := i + 1
		for ; j < len(s) && s[j] != ')'; j++ {
		}
		val, ok := knowledgeMap[s[i+1:j]]
		if ok {
			builder.WriteString(val)
		} else {
			builder.WriteRune('?')
		}
		i = j
	}
	return builder.String()
}

func BalancedString(s string) int {
	counter := map[rune]int{}
	for _, c := range s {
		counter[c]++
	}
	n := len(s)
	partial := n / 4
	balanced := func() bool {
		if counter['Q'] > partial ||
			counter['W'] > partial ||
			counter['E'] > partial ||
			counter['R'] > partial {
			return false
		}
		return true
	}
	if balanced() {
		return 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	r := 0
	res := n
	for l, c := range s {
		for r < n && !balanced() {
			counter[rune(s[r])]--
			r += 1
		}
		if !balanced() {
			break
		}
		res = min(res, r-l)
		counter[c]++
	}
	return res
}

var letterTable = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

var combinations []string

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := digits[index]
		letters := letterTable[rune(digit)]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

var mounthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aa, la := parseDay(arriveAlice), parseDay(leaveAlice)
	ab, lb := parseDay(arriveBob), parseDay(leaveBob)

	if ab > la || aa > lb {
		return 0
	}

	var l, r int
	if aa >= ab {
		l = aa
	} else {
		l = ab
	}
	if la <= lb {
		r = la
	} else {
		r = lb
	}
	return r - l + 1
}

func parseDay(s string) int {
	ss := strings.Split(s, "-")
	b, _ := strconv.Atoi(ss[0])
	e, _ := strconv.Atoi(ss[1])

	for i := b - 2; i >= 0; i-- {
		e += mounthDays[i]
	}
	return e
}

// AddString return the sum of num1 and
// num2, also represented as a string.
func AddString(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}

	sum := ""
	k := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || k != 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}

		s := a + b + k
		k = s / 10
		sum = strconv.Itoa(s%10) + sum
	}

	return sum
}

func SimplifyPath(path string) string {
	names := strings.Split(path, "/")
	p := make([]string, 0, len(names))

	for _, v := range names {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(p) == 0 {
				continue
			}
			p = p[:len(p)-1]
		default:
			p = append(p, v)
		}
	}
	return "/" + strings.Join(p, "/")
}

// SplitWordsBySeparator splits the words in the input slice by the given
// separator character.
// It returns a new slice containing the split words.
// If a word is empty after splitting, it is not included in the result.
// Example usage:
//
//	words := []string{"one.two.three", "four.five", "six"}
//	separator := '.'
//	result := SplitWordsBySeparator(words, separator)
//
//	// result will be []string{"one", "two", "three", "four", "five", "six"}
func SplitWordsBySeparator(words []string, separator byte) []string {
	res := make([]string, 0, len(words))
	for _, word := range words {
		newWord := &strings.Builder{}
		for _, char := range word {
			if byte(char) == separator {
				if newWord.Len() != 0 {
					res = append(res, newWord.String())
				}
				newWord.Reset()
			} else {
				newWord.WriteByte(byte(char))
			}
		}
		if newWord.Len() != 0 {
			res = append(res, newWord.String())
		}
	}
	return res
}
