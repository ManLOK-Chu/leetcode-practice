package main

import (
	"container/list"
	"fmt"
)

func isPalindrome(s string) bool {
	queue := list.New()
	var length = len(s)
	for i := 0; i < length; i++ {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') {
			queue.PushBack(s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			queue.PushBack(s[i] + 32)
		}
	}
	var fV, bV byte
	for queue.Len() > 1 {
		fV = queue.Remove(queue.Front()).(byte)
		bV = queue.Remove(queue.Back()).(byte)
		if fV != bV {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}
