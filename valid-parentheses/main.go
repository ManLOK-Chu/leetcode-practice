package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	item   []byte
	length int
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Push(char byte) {
	s.lazyInit()
	s.item = append(s.item, char)
	s.length++
}

func (s *Stack) Pop() byte {
	if s.length == 0 {
		return 0
	}
	s.length--
	item := s.item[s.length]
	s.item = s.item[0:s.length]
	return item
}

func (s *Stack) lazyInit() {
	if s.item == nil {
		s.item = make([]byte, 0, 32)
	}
}

var tokenMapping = map[byte]byte{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := list.New()
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack.PushBack(s[i])
		case ')', '}', ']':
			if stack.Len() > 0 {
				character := stack.Remove(stack.Back()).(byte)
				if tokenMapping[s[i]] != character {
					return false
				}
			} else {
				return false
			}

		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("))"))
}

//func isValid(s string) bool {
//	var length = len(s)
//	if length%2 == 1 { //奇数长度直接返回false
//		return false
//	}
//	stack := new(Stack)
//	for i := range s {
//		switch s[i] {
//		case '(', '{', '[':
//			stack.Push(s[i])
//		case ')', '}', ']':
//			if ch := stack.Pop(); ch == 0 || ch != tokenMapping[s[i]] {
//				return false
//			}
//		}
//	}
//	if stack.Size() != 0 {
//		return false
//	}
//	return true
//}
