package fib

import (
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "CASE 1", args: args{0}, want: 0},
		{name: "CASE 2", args: args{1}, want: 1},
		{name: "CASE 3", args: args{10}, want: 55},
		{name: "CASE 4", args: args{11}, want: 89},
		{name: "CASE 5", args: args{12}, want: 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.x); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"CASE 1", args{1, 2}, 3}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Fibo25_1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibSlow(28)
	}
}

func Benchmark_Fibo25_2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(28)
	}
}
