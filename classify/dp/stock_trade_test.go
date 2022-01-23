package dp

import "testing"

func Test_maxProfit(t *testing.T) {
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
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "example2",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit2(t *testing.T) {
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
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "example2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit3_(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3_(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit3_() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit4_(t *testing.T) {
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
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit4_(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit4_() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit5_(t *testing.T) {
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
			args: args{prices: []int{1, 2, 3, 0, 2}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5_(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit5_() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit6_(t *testing.T) {
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
			if got := maxProfit6_(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit6_() = %v, want %v", got, tt.want)
			}
		})
	}
}
