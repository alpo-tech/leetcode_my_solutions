package leetcode

import (
	"reflect"
	"testing"
)

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

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{[]int{5, 10, -5}}, want: []int{5, 10}},
		{args: args{[]int{8, -8}}, want: []int{}},
		{args: args{[]int{10, 2, -5}}, want: []int{10}},
		{args: args{[]int{-2, -1, 1, 2}}, want: []int{-2, -1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
