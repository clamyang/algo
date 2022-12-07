package backtracking

import (
	"fmt"
	"testing"
)

func TestCombinationIII(t *testing.T) {
	OptCombinationIII(3, 9, 1)
	fmt.Println(finalRes)
}

func TestCombinationCp(t *testing.T) {
	CombinationCp()
}
