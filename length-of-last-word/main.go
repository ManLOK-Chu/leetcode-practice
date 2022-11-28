package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := len(s)
	var m, n int
	for m = length - 1; m > 0; m-- {
		if s[m] != ' ' {
			break
		}
	}
	for n = m - 1; n >= 0; n-- {
		if s[n] != ' ' {
			continue
		}
		break
	}
	return m - n
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord(" a"))
	fmt.Println(lengthOfLastWord("day"))
}
