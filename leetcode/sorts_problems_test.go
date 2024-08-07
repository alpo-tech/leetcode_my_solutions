package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{nums: []int{5, 2, 3, 1}}, want: []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllRecipes(t *testing.T) {
	type args struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
	}
	tests := []struct {
		args args
		want []string
	}{
		//{args: args{[]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour"}}, want: []string{"bread"}},
		{args: args{[]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}}, want: []string{"bread", "sandwich"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findAllRecipes(tt.args.recipes, tt.args.ingredients, tt.args.supplies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllRecipes() = %v, want %v", got, tt.want)
			}
		})
	}
}
