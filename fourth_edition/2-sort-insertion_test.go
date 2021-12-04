package fourth_edition

import "testing"

func TestInsertion_Sort(t *testing.T) {
	type args struct {
		list []Comparable
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insertion sort",
			args: args{list: sortList},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := BaseSort{SortImpl: Insertion{}}
			sorter.Sort(tt.args.list)
			sorter.Show(tt.args.list)
			if !sorter.IsSorted(tt.args.list) {
				t.Error("not sorted", tt.args.list)
			}
		})
	}
}
