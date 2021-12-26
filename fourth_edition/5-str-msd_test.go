package fourth_edition

import (
	"reflect"
	"testing"
)

func TestMSD_Sort(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name  string
		args  args
		check func(got []string) bool
	}{
		{
			name: "empty",
			args: args{a: nil},
			check: func(got []string) bool {
				return got == nil
			},
		},
		{
			name: "normal",
			args: args{a: []string{"DHFH", "ABCD", "GHI", "GIH", "ACD", "ABH", "ABC", "DHFHABC"}},
			check: func(got []string) bool {
				want := []string{"ABC", "ABCD", "ABH", "ACD", "DHFH", "DHFHABC", "GHI", "GIH"}
				return reflect.DeepEqual(got, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMSD()
			m.Sort(tt.args.a)
			if !tt.check(tt.args.a) {
				t.Errorf("failed, name: %s, list: %v", tt.name, tt.args.a)
			}
		})
	}
}
