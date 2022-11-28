package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var min = 200
	var minIndex int
	for i := range strs {
		n := len(strs[i])
		if n < min {
			minIndex = i
			min = n
		}
	}
	var result = make([]byte, 0, min)
search:
	for i := range strs[minIndex] {
		var ch = strs[minIndex][i]
		for j := range strs {
			if strs[j][i] != ch {
				break search
			}
		}
		result = append(result, ch)
	}
	return string(result)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
