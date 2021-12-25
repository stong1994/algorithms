package fourth_edition

import (
	"reflect"
	"testing"
)

func TestKeyIdxCount_sort(t *testing.T) {
	type args struct {
		a []int
		r int
	}
	tests := []struct {
		name  string
		args  args
		check func(after []int) bool
	}{
		{
			name: "empty",
			args: args{
				a: nil,
				r: 0,
			},
			check: func(after []int) bool {
				return after == nil
			},
		},
		{
			name: "simple",
			args: args{
				a: []int{2, 1, 1, 5, 3, 3},
				r: 5,
			},
			check: func(after []int) bool {
				r := []int{1, 1, 2, 3, 3, 5}
				return reflect.DeepEqual(after, r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyIdxCount{}
			k.sort(tt.args.a, tt.args.r)
			if !tt.check(tt.args.a) {
				t.Errorf("name:%s, a:%v", tt.name, tt.args.a)
			}
		})
	}
}
