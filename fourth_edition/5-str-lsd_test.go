package fourth_edition

import (
	"reflect"
	"testing"
)

func TestLSD_sort(t *testing.T) {
	type args struct {
		a []string
		w int
	}
	tests := []struct {
		name  string
		args  args
		check func(got []string) bool
	}{
		{
			name: "empty",
			args: args{
				a: nil,
				w: 0,
			},
			check: func(got []string) bool {
				return got == nil
			},
		},
		{
			name: "normal",
			args: args{
				a: []string{"DHF", "ABC", "GHI", "GIH", "ACD", "ABH"},
				w: 3,
			},
			check: func(got []string) bool {
				want := []string{"ABC", "ABH", "ACD", "DHF", "GHI", "GIH"}
				return reflect.DeepEqual(got, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LSD{}
			l.sort(tt.args.a, tt.args.w)
			if !tt.check(tt.args.a) {
				t.Errorf("failed, name: %s, list: %v", tt.name, tt.args.a)
			}
		})
	}
}
