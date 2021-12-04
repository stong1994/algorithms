package fourth_edition

import "testing"

func TestSelection_Sort(t *testing.T) {
	type args struct {
		list []Comparable
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal test",
			args: args{list: []Comparable{2, 4, 1, 2, 9, 11, 5, 6, 3, 1, 0, 10, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := BaseSort{SortImpl: Selection{}}
			sorter.Sort(tt.args.list)
			sorter.Show(tt.args.list)
			if !sorter.IsSorted(tt.args.list) {
				t.Error("not sorted", tt.args.list)
			}
		})
	}
}
