package classify

import "testing"

func Test_shortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example1",
			grid: [][]int{{0, 1}, {1, 0}},
			want: 2,
		},
		{
			name: "example2",
			grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			want: 4,
		},
		{
			name: "example3",
			grid: [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			want: -1,
		},
		/*
			0, 1, 1, 0, 0, 0
			0, 1, 0, 1, 1, 0
			0, 1, 1, 0, 1, 0
			0, 0, 0, 1, 1, 0
			1, 1, 1, 1, 1, 0
			1, 1, 1, 1, 1, 0
		*/
		{
			name: "example4",
			grid: [][]int{{0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}},
			want: 14,
		},
		/*
			0, 1, 0, 0, 0, 0
			0, 1, 0, 1, 1, 0
			0, 1, 1, 0, 1, 0
			0, 0, 0, 0, 1, 0
			1, 1, 1, 1, 1, 0
			1, 1, 1, 1, 1, 0
		*/
		{
			name: "example5",
			grid: [][]int{{0, 1, 0, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    12,
			want: 3,
		},
		{
			name: "example2",
			n:    13,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
			if got := numSquares2(tt.n); got != tt.want {
				t.Errorf("numSquares2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
