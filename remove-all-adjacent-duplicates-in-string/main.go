package main

import (
	"container/list"
	"fmt"
)

func removeDuplicates(s string) string {
	stack := list.New()
	var length = len(s)
	var n int
	for i := 0; i < length; i++ {
		if val := stack.Back(); val != nil {
			if val.Value.(byte) == s[i] {
				n--
				stack.Remove(val)
				continue
			}
		}
		n++
		stack.PushBack(s[i])
	}
	var result = make([]byte, n)
	for n > 0 {
		result[n-1] = stack.Remove(stack.Back()).(byte)
		n--
	}
	return string(result)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}
