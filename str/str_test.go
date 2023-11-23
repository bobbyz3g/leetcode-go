package str

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	var tests567 = []struct {
		a      string
		b      string
		result bool
	}{
		{
			a:      "ab",
			b:      "eidbaooo",
			result: true,
		},
		{
			a:      "abc",
			b:      "b",
			result: false,
		},
	}
	for _, v := range tests567 {
		assert.Equal(t, CheckInclusion(v.a, v.b), v.result)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	var tests14 = []struct {
		strs   []string
		output string
	}{
		{
			strs:   []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			strs:   []string{"a"},
			output: "a",
		},
	}

	for _, v := range tests14 {
		assert.Equal(t, LongestCommonPrefix(v.strs), v.output)
	}
}

func TestPartition(t *testing.T) {
	var tests131 = []struct {
		s   string
		out [][]string
	}{
		{
			s:   "aab",
			out: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:   "a",
			out: [][]string{{"a"}},
		},
	}

	for _, v := range tests131 {
		assert.Equal(t, v.out, Partition(v.s))
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		s   string
		out int
	}{
		{
			s:   "abcabcbb",
			out: 3,
		},
		{
			s:   "bbbbb",
			out: 1,
		},
		{
			s:   "pwwkew",
			out: 3,
		},
		{
			s:   "a",
			out: 1,
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.out, LengthOfLongestSubstring(v.s))
	}
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{name: "empty", str: "", want: 0},
		{name: "one space", str: " ", want: 0},
		{name: "two space", str: " ", want: 0},
		{name: "only space", str: "    ", want: 0},
		{name: "one rune", str: "a", want: 1},
		{name: "two rune", str: "ab", want: 2},
		{name: "two rune with space", str: "ab ab", want: 2},
		{name: "end with space", str: "ab ", want: 2},
		{name: "normal", str: "Hello World", want: 5},
		{name: "   fly me   to   the moon  ", str: "   fly me   to   the moon  ", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LengthOfLastWord(tt.str), "LengthOfLastWord(%v)", tt.str)
		})
	}
}

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "11, 1", args: args{a: "11", b: "1"}, want: "100"},
		{name: "1010, 1011", args: args{a: "1010", b: "1011"}, want: "10101"},
		{name: "0, 0", args: args{a: "0", b: "0"}, want: "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddBinary(tt.args.a, tt.args.b), "AddBinary(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "1"},
			want: true,
		},
		{
			name: "race a car",
			args: args{s: "race a cara"},
			want: false,
		},
		{
			name: "A man, a plan, a canal: Panama",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "empty string",
			args: args{"s"},
			want: true,
		},
		{
			name: "OP",
			args: args{"0P"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsPalindrome(tt.args.s), "IsPalindrome(%v)", tt.args.s)
		})
	}
}

func TestConvertToTitle(t *testing.T) {
	type args struct {
		column int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{column: 1},
			want: "A",
		},
		{
			name: "Z",
			args: args{column: 26},
			want: "Z",
		},
		{
			name: "ZY",
			args: args{column: 701},
			want: "ZY",
		},
		{
			name: "FXSHRXW",
			args: args{column: 2147483647},
			want: "FXSHRXW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ConvertToTitle(tt.args.column), "ConvertToTitle(%v)", tt.args.column)
		})
	}
}

func TestTitleToNumber(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{title: "A"},
			want: 1,
		},
		{
			name: "Z",
			args: args{title: "Z"},
			want: 26,
		},
		{
			name: "ZY",
			args: args{title: "ZY"},
			want: 701,
		},
		{
			name: "FXSHRXW",
			args: args{title: "FXSHRXW"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TitleToNumber(tt.args.title), "TitleToNumber(%v)", tt.args.title)
		})
	}
}

func TestIsIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "egg-add",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "foo-bar",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "paper-title",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsIsomorphic(tt.args.s, tt.args.t), "IsIsomorphic(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}

func TestCheckIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abcdd",
			args: args{sentence: "abcdd"},
			want: false,
		},
		{
			name: "thequickbrownfoxjumpsoverthelazydog",
			args: args{sentence: "thequickbrownfoxjumpsoverthelazydog"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CheckIfPangram(tt.args.sentence), "CheckIfPangram(%v)", tt.args.sentence)
		})
	}
}

func TestNumbersAscending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			args: args{
				s: "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			},
			want: true,
		},
		{
			name: "hello world 5 x 5",
			args: args{
				s: "hello world 5 x 5",
			},
			want: false,
		},
		{
			name: "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
			args: args{
				s: "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
			},
			want: false,
		},
		{
			name: "4 5 11 26",
			args: args{
				s: "4, 5, 11, 26",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NumbersAscending(tt.args.s), "NumbersAscending(%v)", tt.args.s)
		})
	}
}

func TestDigitCount(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1210",
			args: args{num: "1210"},
			want: true,
		},
		{
			name: "030",
			args: args{num: "030"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DigitCount(tt.args.num), "DigitCount(%v)", tt.args.num)
		})
	}
}

func TestEvaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "(name)is(age)yearsold",
			args: args{
				s: "(name)is(age)yearsold",
				knowledge: [][]string{
					{"name", "bob"},
					{"age", "two"},
				},
			},
			want: "bobistwoyearsold",
		},
		{
			name: "hi(name)",
			args: args{
				s: "hi(name)",
				knowledge: [][]string{
					{"a", "b"},
				},
			},
			want: "hi?",
		},
		{
			name: "(a)(a)(a)aaa",
			args: args{
				s: "(a)(a)(a)aaa",
				knowledge: [][]string{
					{"a", "yes"},
				},
			},
			want: "yesyesyesaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Evaluate(tt.args.s, tt.args.knowledge), "Evaluate(%v, %v)", tt.args.s, tt.args.knowledge)
		})
	}
}

func TestBalancedString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "QWER",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				s: "QQWE",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				s: "QQQW",
			},
			want: 2,
		},
		{
			name: "case4",
			args: args{
				s: "QQQQ",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BalancedString(tt.args.s), "BalancedString(%v)", tt.args.s)
		})
	}
}

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "case2",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "case3",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterCombinations(tt.args.digits)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			assert.Equalf(t, tt.want, got, "LetterCombinations(%v)", tt.args.digits)
		})
	}
}

func TestCountDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arriveAlice: "08-15",
				leaveAlice:  "08-18",
				arriveBob:   "08-16",
				leaveBob:    "08-19",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				arriveAlice: "10-01",
				leaveAlice:  "10-31",
				arriveBob:   "11-01",
				leaveBob:    "12-31",
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				arriveAlice: "08-15",
				leaveAlice:  "08-18",
				arriveBob:   "08-14",
				leaveBob:    "08-20",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob), "CountDaysTogether(%v, %v, %v, %v)", tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob)
		})
	}
}

func TestAddString(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "5",
		},
		{
			name: "case2",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "case3",
			args: args{
				num1: "123",
				num2: "0",
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddString(tt.args.num1, tt.args.num2), "AddString(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
