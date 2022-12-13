package main

import "fmt"

func replaceSpace(s string) string {
	var spaceCnt, length = 0, len(s)
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			spaceCnt++
		}
	}
	if spaceCnt == 0 {
		return s
	}
	var newLength = length + 2*spaceCnt
	var result = make([]byte, newLength)
	var n, old = newLength - 1, length - 1
	for n >= 0 {
		result[n] = s[old]
		if s[old] == ' ' {
			result[n] = '0'
			result[n-1] = '2'
			result[n-2] = '%'
			n -= 2
		}
		n--
		old--
	}
	return string(result)
}

func main() {
	//fmt.Println(replaceSpace(" "))
	fmt.Println(replaceSpace("We are happy.")) //We%20are%20happy.
}
