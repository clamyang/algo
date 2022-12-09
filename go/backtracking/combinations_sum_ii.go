package backtracking

// 040 组合总和II

import "fmt"

func CombinationSumII() {
	candidates := []int{1, 1, 2} // 排序后的有序数组
	target := 3
	used := make([]bool, len(candidates))
	data := []int{}

	recursiveSumII(candidates, target, 0, &used, &data)
}

func recursiveSumII(candidates []int, target, startIdx int, used *[]bool, data *[]int) {
	if sum(*data) == target {
		fmt.Println(*data)
		return
	} else if sum(*data) > target {
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !((*used)[i-1]) {
			continue
		}

		// 利用 used 去重
		*data = append(*data, candidates[i])
		(*used)[i] = true
		recursiveSumII(candidates, target, i+1, used, data)
		(*used)[i] = false
		*data = (*data)[:len(*data)-1]
	}
}
