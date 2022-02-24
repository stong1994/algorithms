package dp

import "testing"

func Test_numberOfGoodSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 6,
		},
		{
			name: "example2",
			args: args{nums: []int{4, 2, 3, 15}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfGoodSubsets(tt.args.nums); got != tt.want {
				t.Errorf("numberOfGoodSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        11,
			},
			want: false,
		},
		{
			name: "example2",
			args: args{
				maxChoosableInteger: 10,
				desiredTotal:        40,
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				maxChoosableInteger: 4,
				desiredTotal:        6,
			},
			want: true,
		},
		{
			name: "example4",
			args: args{
				maxChoosableInteger: 5,
				desiredTotal:        50,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("canIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{matchsticks: []int{1, 1, 2, 2, 2}},
			want: true,
		},
		{
			name: "example2",
			args: args{matchsticks: []int{3, 3, 3, 3, 4}},
			want: false,
		},
		{
			name: "example3",
			args: args{matchsticks: []int{3, 3, 3, 3, 4, 4, 4, 2, 2}},
			want: true,
		},
		{
			name: "example4",
			args: args{matchsticks: []int{2, 2, 2, 2, 4, 4, 4, 4, 4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare3(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countArrangement(t *testing.T) {
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
			args: args{n: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangement2(tt.args.n); got != tt.want {
				t.Errorf("countArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
