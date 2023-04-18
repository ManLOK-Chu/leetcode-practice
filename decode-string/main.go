package main

import (
	"fmt"
	"strconv"
	"strings"
)

//给定一个经过编码的字符串，返回它解码后的字符串。
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。
func decodeString(s string) string {
	var stk []string
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits, offset := getDigits(s, ptr)
			ptr += offset
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr int) (string, int) {
	var result []byte
	i := ptr
	for ; s[i] >= '0' && s[i] <= '9'; i++ {
		result = append(result, s[i])
	}
	return string(result), i - ptr
}

func getString(v []string) string {
	return strings.Join(v, "")
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("abc3[cd]xyz"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("3[a2[c]]"))
}
