package backtracking

func partition(s string) [][]string {
	var tmpString []string //切割字符串集合
	var res [][]string     //结果集合
	backTracking(s, tmpString, 0, &res)
	return res
}
func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) { //到达字符串末尾了
		//进行一次切片拷贝，怕之后的操作影响tmpString切片内的值
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}
	for i := startIndex; i < len(s); i++ {
		//处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，若为回文，则加入到tmpString，否则继续后移，找到回文区间）（这里为一层处理）
		if isPalindrome(s) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		//递归
		backTracking(s, tmpString, i+1, res)
		//回溯
		tmpString = tmpString[:len(tmpString)-1]
	}
}

func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i-1, j+1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
