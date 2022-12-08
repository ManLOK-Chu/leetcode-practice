package main

import "fmt"

const m = uint32(1)

func hammingWeight(num uint32) int {
	var count int
	fmt.Printf("%032b\n", num)
	for num != 0 {
		if num&m == m {
			count++
		}
		num = num >> 1
		fmt.Printf("%032b\n", num)
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(13))
	fmt.Println(hammingWeight(200))
	//fmt.Println(hammingWeight(37777777775))
}
