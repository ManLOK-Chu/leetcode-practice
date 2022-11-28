package main

import "fmt"

type Stack struct {
	item []byte
}

func (s *Stack) Size() int {
	return len(s.item)
}

func (s *Stack) Push(char byte) {
	s.item = append(s.item, char)
}

func (s *Stack) Pop() byte {
	length := s.Size()
	if length == 0 {
		return 0
	}
	item := s.item[length-1]
	s.item = s.item[0 : length-1]
	return item
}

var tokenMapping = map[byte]byte{
	'}': '{',
	')': '(',
	']': '[',
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
func isValid(s string) bool {
	var length = len(s)
	if length%2 == 1 { //奇数长度直接返回false
		return false
	}
	stack := new(Stack)
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack.Push(s[i])
		case ')', '}', ']':
			if ch := stack.Pop(); ch == 0 || ch != tokenMapping[s[i]] {
				return false
			}
		}
	}
	if stack.Size() != 0 {
		return false
	}
	return true
}
