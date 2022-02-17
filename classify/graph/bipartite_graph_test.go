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
