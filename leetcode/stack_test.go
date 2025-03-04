package leetcode

import "testing"

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{"abb*cdfg*****x*"}, want: "a"},
		{args: args{"leet**cod*e"}, want: "lecoe"},
		{args: args{"erase*****"}, want: ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
