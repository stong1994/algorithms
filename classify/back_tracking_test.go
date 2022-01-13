package classify

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "23",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "empty",
			digits: "",
			want:   []string{},
		},
		{
			name:   "2",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "25525511135",
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "0000",
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			name: "1111",
			s:    "1111",
			want: []string{"1.1.1.1"},
		},
		{
			name: "010010",
			s:    "010010",
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name: "101023",
			s:    "101023",
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCCED",
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "SEE",
			},
			want: true,
		},
		{
			name: "example3",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCB",
			},
			want: false,
		},
		{
			name: "example4",
			args: args{
				board: [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}},
				word:  "aaaaaaaaaaaaa",
			},
			want: false,
		},
		{
			name: "example5",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCESEEEFS",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTreePaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []string
	}{
		{
			name: "example1",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   3,
					Right: nil,
					Left:  nil,
				},
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   5,
						Right: nil,
						Left:  nil,
					},
				},
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "example2",
			root: &TreeNode{
				Val: 1,
			},
			want: []string{"1"},
		},
		{
			name: "example3",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: []string{"6->1->3->2", "6->1->3->5->4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example1",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "example2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "example3",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
			if got := permuteOpt(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example1",
			nums: []int{1, 1, 2},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
			//want: [][]int{{1, 1, 2},1, 1, 2}, {1, 2, 1},{1, 2, 1}, {2, 1, 1}{2, 1, 1}},
		},
		{
			name: "example2",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.nums)
			//sort.Slice(got, func(i, j int) bool {
			//	return less(got[i], got[j])
			//})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(a, b []int) bool {
	for i := range a {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}},
		},
		{
			name: "example2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
			if got := combineOpt(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}
