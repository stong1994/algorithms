package tree

import (
	"reflect"
	"testing"
)

func Test_pmt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{str: "abcaabcaad"},
			want: []int{0, 0, 0, 1, 1, 2, 3, 4, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pmt(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
