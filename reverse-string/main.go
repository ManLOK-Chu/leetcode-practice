package main

import "fmt"

func reverseString(s []byte) {
	var length = len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func main() {
	var s = []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
	s = []byte("Hannah")
	reverseString(s)
	fmt.Println(string(s))
}
