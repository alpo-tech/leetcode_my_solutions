package leetcode

import (
	"testing"
)

func TestConvertToTitle_168(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{1}, want: "A"},    // TODO: Add test cases.
		{args: args{27}, want: "AA"},  // TODO: Add test cases.
		{args: args{28}, want: "AB"},  // TODO: Add test cases.
		{args: args{701}, want: "ZY"}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ConvertToTitle_168(tt.args.columnNumber); got != tt.want {
				t.Errorf("ConvertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
