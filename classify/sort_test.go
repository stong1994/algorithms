package classify

import "testing"

func Test_topN(t *testing.T) {
	var sortList = []int{2, 4, 1, 2, 9, 11, 5, 6, 3, 1, 0, 10, 8, -1}

	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "topN", args: args{sortList, 5}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topN(tt.args.list, tt.args.k); got != tt.want {
				t.Errorf("topN() = %v, want %v", got, tt.want)
			}
		})
	}
}
