package graph

import "testing"

func Test_isBipartite_uf(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite_uf(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite_uf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				n: 4,
				dislikes: [][]int{
					{1, 2},
					{1, 3},
					{2, 4},
				},
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				n: 4,
				dislikes: [][]int{
					{1, 2},
					{1, 3},
					{2, 3},
				},
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				n: 5,
				dislikes: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
					{1, 5},
				},
			},
			want: false,
		},
		{
			name: "example4",
			args: args{
				n: 5,
				dislikes: [][]int{
					{1, 2},
					{1, 3},
					{1, 4},
					{1, 5},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
			if got := possibleBipartition_dfs(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition_dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
