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

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "example2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "example3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "example4",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsNormal(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsNormal() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsMemOpt(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsMemOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example1",
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			name: "example2",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
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
