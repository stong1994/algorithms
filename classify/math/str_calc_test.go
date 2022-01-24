package math

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "example2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "example3",
			args: args{
				a: "0",
				b: "0",
			},
			want: "0",
		},
		{
			name: "example4",
			args: args{
				a: "0",
				b: "1",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		},
		{
			name: "example2",
			args: args{
				num1: "456",
				num2: "77",
			},
			want: "533",
		},
		{
			name: "example3",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
