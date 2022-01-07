package classify

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "4",
			x:    4,
			want: 2,
		},
		{
			name: "8",
			x:    8,
			want: 2,
		},
		{
			name: "1",
			x:    1,
			want: 1,
		},
		{
			name: "2",
			x:    2,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
