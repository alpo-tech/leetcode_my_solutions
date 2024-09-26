package main

import "testing"

type test int

func TestIntMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{10, 5}, want: 5},
		{args: args{10, 15}, want: 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IntMin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
