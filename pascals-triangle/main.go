package main

import "fmt"

func generate(numRows int) [][]int {
	var result = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if i > 1 && j > 0 && i > j {
				result[i][j] =
					result[i-1][j] +
						result[i-1][j-1]
			} else {
				result[i][j] = 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}
