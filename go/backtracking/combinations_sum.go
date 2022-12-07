package backtracking

import "fmt"

/*
	为什么需要startIndex来控制for循环的起始位置，对于组合问题，什么时候需要startIndex呢？
	1.如果是一个集合来求组合的话，就需要startIndex
	2.如果是多个集合取组合，各个集合之间相互不影响，那么就不用startIndex，例如：17.电话号码的字母组合
	注意以上我只是说求组合的情况，如果是排列问题，又是另一套分析的套路。
*/

func CombinationSum() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := []int{}

	// recursiveSum(candidates, target, &result)
	recursiveSumOpt(candidates, target, &result, 0)
}

// 该种方式输出如下
/*
	[2 2 3]
	[2 3 2]
	[3 2 2]
	[7]
*/
func recursiveSum(candidates []int, target int, result *[]int) {
	if sum(*result) == target {
		fmt.Println(*result)
		return
	} else if sum(*result) > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		*result = append(*result, candidates[i])
		recursiveSum(candidates, target, result)
		*result = (*result)[:len(*result)-1]
	}
}

func recursiveSumOpt(candidates []int, target int, result *[]int, startIdx int) {
	if sum(*result) == target {
		fmt.Println(*result)
		return
	} else if sum(*result) > target {
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		*result = append(*result, candidates[i])
		recursiveSumOpt(candidates, target, result, i)
		*result = (*result)[:len(*result)-1]
	}
}
