package backtracking

import (
	"fmt"
	"testing"
)

func TestCombination(t *testing.T) {
	type args struct {
		n          int
		k          int
		startIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"n = 4, k = 2", args{4, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Combination(tt.args.n, tt.args.k, tt.args.startIndex)
		})
	}

	fmt.Println(result)

	for idx := range result {
		fmt.Printf("%p\n", &result[idx])
	}
}
