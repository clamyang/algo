package backtracking

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
