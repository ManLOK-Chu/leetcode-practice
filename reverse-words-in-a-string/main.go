package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var left, right, length = 0, 0, len(s)
	var results []string
	for left < length {
		if s[left] == ' ' {
			left++
			continue
		}
		for right = left + 1; right < length && s[right] != ' '; right++ {
		}
		results = append(results, s[left:right])
		left = right
	}
	for i, l := 0, len(results); i < l/2; i++ {
		results[i], results[l-i-1] = results[l-i-1], results[i]
	}
	return strings.Join(results, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
