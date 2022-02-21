package foreach

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{dominoes: "RR.L"},
			want: "RR.L",
		},
		{
			name: "example2",
			args: args{dominoes: ".L.R...LR..L.."},
			want: "LL.RR.LLRRLL..",
		},
		{
			name: "example3",
			args: args{dominoes: ".L.R."},
			want: "LL.RR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
