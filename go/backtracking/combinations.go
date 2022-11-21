package backtracking

// #77 combination 组合

var (
	result [][]int
	path   []int
)

func Combination(n, k, startIndex int) {
	if len(path) == k {
		data := make([]int, len(path))
		copy(data, path)
		result = append(result, data)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		Combination(n, k, i+1)
		path = path[:len(path)-1]
	}
}
