package leetcode

import (
	"reflect"
	"testing"
)

func TestCountNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNodes_222(tt.args.root); got != tt.want {
				t.Errorf("CountNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		args args
		want []string
	}{}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := BinaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
