package main

import "fmt"

var letterMap = map[byte]string{
	'2': "abc",  // 2
	'3': "def",  // 3
	'4': "ghi",  // 4
	'5': "jkl",  // 5
	'6': "mno",  // 6
	'7': "pqrs", // 7
	'8': "tuv",  // 8
	'9': "wxyz", // 9
}

var result []string
var bs []byte

func letterCombinations(digits string) []string {
	result = nil
	bs = nil
	if len(digits) == 0 {
		return nil
	}
	backtracking([]byte(digits), 0)
	return result
}

func backtracking(digits []byte, index int) {
	if index == len(digits) {
		result = append(result, string(bs))
		return
	}
	digit := digits[index]
	letters := letterMap[digit]
	for i := 0; i < len(letters); i++ {
		bs = append(bs, letters[i])
		backtracking(digits, index+1)
		bs = bs[:len(bs)-1]
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("1*#"))
}
