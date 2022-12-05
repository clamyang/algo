package backtracking

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 || length > 4 {
		return nil
	}
	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	res := make([]string, 0)
	recursion("", digits, 0, digitsMap, &res)
	return res
}
func recursion(tempString, digits string, Index int, digitsMap [10]string, res *[]string) { //index表示第几个数字
	if len(tempString) == len(digits) { //终止条件，字符串长度等于digits的长度
		*res = append(*res, tempString)
		return
	}
	tmpK := digits[Index] - '0' // 将index指向的数字转为int（确定下一个数字）
	letter := digitsMap[tmpK]   // 取数字对应的字符集
	for i := 0; i < len(letter); i++ {
		tempString = tempString + string(letter[i]) //拼接结果
		recursion(tempString, digits, Index+1, digitsMap, res)
		tempString = tempString[:len(tempString)-1] //回溯
	}
}

func LetterPhoneNumber() {
	number, length := "23", len("23")

	maps := map[byte]string{
		byte(2): "abc",  // 2
		byte(3): "def",  // 3
		byte(4): "ghi",  // 4
		byte(5): "jkl",  // 5
		byte(6): "mno",  // 6
		byte(7): "pqrs", // 7
		byte(8): "tuv",  // 8
		byte(9): "wxyz", // 9
	}

	recursiveLetter(number, length, &[]string{}, 0, maps)
}

func recursiveLetter(number string, length int, data *[]string, index int, maps map[byte]string) {
	if index == length {
		fmt.Println(*data)
		return
	}


	numbers := maps[getNumberfromIndex(index, number)]

	for i := 0; i < len(numbers); i++ {
		*data = append(*data, string(numbers[i]))
		recursiveLetter(number, length, data, index+1, maps)
		*data = (*data)[:len(*data)-1]
	}
}

func getNumberfromIndex(index int, str string) byte {
	return str[index]
}
