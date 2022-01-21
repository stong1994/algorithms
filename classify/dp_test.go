package classify

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    2,
			want: 2,
		},
		{
			name: "example2",
			n:    3,
			want: 3,
		},
		{
			name: "example3",
			n:    4,
			want: 5,
		},
		{
			name: "example4",
			n:    5,
			want: 8,
		},
		{
			name: "example5",
			n:    6,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
			if got := climbStairs_backTrack(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
			if got := climbStairsFib(tt.n); got != tt.want {
				t.Errorf("climbStairsFib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
			if got := robNormal(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
			if got := robNormalOpt(tt.nums); got != tt.want {
				t.Errorf("robNormalOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "example2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.nums); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "example1",
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integerBreak(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    2,
			want: 1,
		},
		{
			name: "example2",
			n:    10,
			want: 36,
		},
		{
			name: "example3",
			n:    8,
			want: 18,
		},
		{
			name: "example4",
			n:    11,
			want: 54,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example1",
			s:    "12",
			want: 2,
		},
		{
			name: "example2",
			s:    "226",
			want: 3,
		},
		{
			name: "example3",
			s:    "0",
			want: 0,
		},
		{
			name: "example4",
			s:    "06",
			want: 0,
		},
		{
			name: "example5",
			s:    "10",
			want: 1,
		},
		{
			name: "example6",
			// [1]
			// [1 1] [11]
			// [1 1 2] [11 2] [1 12]
			// [1 1 2 3] [11 2 3] [1 12 3] [1 1 23] [11 23]
			s:    "1123", // [1 1 2 3] [11 2 3] [1 12 3] [1 1 23] [11 23]
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "example2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "example3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "example4",
			nums: []int{2, 3, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
			if got := lengthOfLISOpt(tt.nums); got != tt.want {
				t.Errorf("lengthOfLISOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLongestChain(t *testing.T) {
	tests := []struct {
		name  string
		pairs [][]int
		want  int
	}{
		{
			name:  "example0",
			pairs: [][]int{{}},
			want:  1,
		},
		{
			name:  "example1",
			pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:  2,
		},
		{
			name:  "example2",
			pairs: [][]int{{1, 2}, {7, 8}, {4, 5}},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wiggleMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{1, 7, 4, 9, 2, 5},
			want: 6,
		},
		{
			name: "example1",
			nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			want: 7,
		},
		{
			name: "example1",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "example",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "example4",
			args: args{
				text1: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				text2: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			want: 210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistance2(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				word1: "sea",
				word2: "eat",
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				word1: "leetcode",
				word2: "etco",
			},
			want: 4,
		}, {
			name: "example3",
			args: args{
				word1: "mart",
				word2: "karma",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance2(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "example2",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "example3",
			args: args{n: 17},
			want: 17,
		},
		{
			name: "example4",
			args: args{n: 12},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
			if got := minSteps2(tt.args.n); got != tt.want {
				t.Errorf("minSteps2() = %v, want %v", got, tt.want)
			}
		})
	}
}
