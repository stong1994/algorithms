package classify

import "testing"

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example1",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "example2",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
		{
			name: "example3",
			nums: []int{1, 1},
			want: true,
		},
		{
			name: "example4",
			nums: []int{14, 9, 8, 4, 3, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
			if got := canPartitionOpt(tt.nums); got != tt.want {
				t.Errorf("canPartitionOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "example3",
			args: args{
				nums:   []int{1, 0},
				target: 1,
			},
			want: 2,
		}, {
			name: "example4",
			args: args{
				nums:   []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
				target: 1,
			},
			want: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			want: 4,
		},
		{
			name: "example2",
			args: args{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "example3",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "example4",
			args: args{
				coins:  []int{1},
				amount: 1,
			},
			want: 1,
		},
		{
			name: "example5",
			args: args{
				coins:  []int{1},
				amount: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 4,
		},
		{
			name: "example2",
			args: args{
				amount: 3,
				coins:  []int{2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
			if got := change2(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "example3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:   []int{1, 2, 3},
				target: 4,
			},
			want: 7,
		},
		{
			name: "example2",
			args: args{
				nums:   []int{9},
				target: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				prices: []int{1, 2, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
			if got := maxProfit4(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit5(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
		{
			name: "example2",
			args: args{
				prices: []int{1, 3, 7, 5, 10, 3},
				fee:    3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit6(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}},
			want: 6,
		},
		{
			name: "example2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "example3",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "example4",
			args: args{prices: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit6() = %v, want %v", got, tt.want)
			}
			if got := maxProfit6Opt(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit6Opt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDiff(t *testing.T) {
	type args struct {
		prices []int
		start  int
		end    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				prices: []int{1, 3, 5},
				start:  0,
				end:    2,
			},
			want: 4,
		},
		{
			name: "example1",
			args: args{
				prices: []int{1, 1, 1},
				start:  0,
				end:    2,
			},
			want: 0,
		},
		{
			name: "example2",
			args: args{
				prices: []int{2, 3, 5, 1, 6},
				start:  0,
				end:    4,
			},
			want: 5,
		},
		{
			name: "example3",
			args: args{
				prices: []int{2, 3, 7, 1, 5},
				start:  0,
				end:    4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.prices, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit7(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
				// 0 -3 0 -3 0
				// 0 -2 0 -2 0
				// 0 -2 4 -2 0
				// 0 -2 4 -1 0
				// 0  0 4  4 0
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit7(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit7() = %v, want %v", got, tt.want)
			}
		})
	}
}
