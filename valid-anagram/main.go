package main

import "fmt"

func isAnagram(s string, t string) bool {
	var m = make([]int, 26)
	var index byte
	for i := 0; i < len(s); i++ {
		index = s[i] - 'a'
		m[index]++
	}
	for i := 0; i < len(t); i++ {
		index = t[i] - 'a'
		if m[index] == 0 {
			return false
		}
		m[index]--
	}
	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
	//var m = make(map[byte]int, 26)
	//for i := range s {
	//	m[s[i]]++
	//}
	//var val int
	//var exist bool
	//for i := range t {
	//	val, exist = m[t[i]]
	//	if !exist { //t中出现了s中没出现过的字符
	//		return false
	//	}
	//	if val > 1 {
	//		m[t[i]] = val - 1
	//	} else {
	//		delete(m, t[i])
	//	}
	//}
	//for range m {
	//	return false
	//}
	//return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
