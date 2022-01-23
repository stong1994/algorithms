package dp

import "testing"

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
