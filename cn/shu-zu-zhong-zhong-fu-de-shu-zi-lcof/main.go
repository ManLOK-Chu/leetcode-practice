package main

import "fmt"

type symbol struct {
}

func findRepeatDocument(documents []int) int {
	var s symbol
	var dict = make([]*symbol, 100000)
	for i := range documents {
		if dict[documents[i]] != nil {
			return documents[i]
		}
		dict[documents[i]] = &s
	}
	return -1
}

func main() {
	fmt.Println(findRepeatDocument([]int{2, 5, 3, 0, 5, 0}))
}
