package backtracking

import "fmt"

// #216 组合总和iii

var (
	finalRes [][]int
	target   []int
)

func CombinationIII(k, n, startIdx int) {
	if len(target) == k && sum(target) == n {
		data := make([]int, len(target))
		copy(data, target)
		finalRes = append(finalRes, data)
		return
	}

	for i := startIdx; i <= 9; i++ {
		target = append(target, i)
		CombinationIII(k, n, i+1)
		target = target[:len(target)-1]
	}
}

func sum(arr []int) int {
	result := 0
	for _, data := range arr {
		result += data
	}
	return result
}

func OptCombinationIII(k, n, startIdx int) {
	if sum(target) > n {
		return
	}

	if len(target) == k && sum(target) == n {
		data := make([]int, len(target))
		copy(data, target)
		finalRes = append(finalRes, data)
		return
	}

	for i := startIdx; i <= 9; i++ {
		target = append(target, i)
		OptCombinationIII(k, n, i+1)
		target = target[:len(target)-1]
	}
}

func CombinationCp() {
	// 在 1-9 中找到合为 9 的所有三个数字的集合
	n := 9
	k := 3
	startIndex := 1

	recursive(n, k, startIndex, &[]int{})
}

func recursive(n, k, startIndex int, data *[]int) {
	if len(*data) == k && sum(*data) == n {
		fmt.Println(data)
		return
	}

	for i := startIndex; i <= n; i++ {
		*data = append(*data, i)
		recursive(n, k, i+1, data)
		*data = (*data)[:len(*data)-1]
	}
}
