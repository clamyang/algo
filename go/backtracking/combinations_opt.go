package backtracking

var (
	optResult [][]int
	optPath   []int
)

// 剪枝，过滤掉不满足条件的情况
// 核心思想：如果从当前位置起始，剩余元素个数不满足我们需要的个数，就没必要继续执行

func OptCombination(n, k, startIndex int) {
	if len(optPath) == k {
		data := make([]int, len(optPath))
		copy(data, optPath)
		optResult = append(optResult, data)
		return
	}

	for i := startIndex; i <= (n - (k - len(optPath)) + 1); i++ {
		optPath = append(optPath, i)
		OptCombination(n, k, i+1)
		optPath = optPath[:len(optPath)-1]
	}
}
