package main

import (
	"container/list"
	"fmt"
)

func addBinary(a string, b string) string {
	var i, j = len(a) - 1, len(b) - 1
	var sum, carry byte
	var stack = list.New()
	for i >= 0 || j >= 0 {
		sum = carry
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		stack.PushBack(sum%2 + '0')
		carry = sum / 2
	}
	if carry != 0 {
		stack.PushBack(carry + '0')
	}
	var result = make([]byte, stack.Len())
	for n := 0; n < len(result); n++ {
		result[n] = stack.Remove(stack.Back()).(byte)
	}
	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("1", "1"))
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
