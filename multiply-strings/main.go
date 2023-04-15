package main

import "fmt"

//注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
//num1,num2 为非负整数
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var i, n = len(num1) - 1, 0
	var res string
	for i >= 0 {
		tmp := multiplySimple(num2, num1[i], n)
		res = addStrings(res, tmp)
		i--
		n++
	}
	return res
}

func multiplySimple(num1 string, num2 byte, padding int) string {
	if num2 == '0' {
		return "0"
	}
	var sum []byte
	var n = len(num1) - 1
	var carry uint8
	var y = num2 - '0'
	for n >= 0 || carry != 0 {
		var x uint8
		if n >= 0 {
			x = num1[n] - '0'
		}
		res := x*y + carry
		carry = res / 10
		res = res % 10
		sum = append(sum, res+'0')
		n--
	}
	for i, j := 0, len(sum)-1; i < j; {
		sum[i], sum[j] = sum[j], sum[i]
		i++
		j--
	}
	for padding > 0 {
		sum = append(sum, '0')
		padding--
	}
	return string(sum)
}

func addStrings(num1 string, num2 string) string {
	var i, j = len(num1) - 1, len(num2) - 1
	var carry uint8
	var sum []byte
	for i >= 0 || j >= 0 || carry != 0 {
		var x, y uint8
		if i >= 0 {
			x = num1[i] - '0'
		}
		if j >= 0 {
			y = num2[j] - '0'
		}
		res := x + y + carry
		carry = res / 10
		res = res % 10
		sum = append(sum, res+'0')
		i--
		j--
	}
	for i, j = 0, len(sum)-1; i < j; {
		sum[i], sum[j] = sum[j], sum[i]
		i++
		j--
	}
	return string(sum)
}

func main() {
	fmt.Println(multiply("1", "456"))
	fmt.Println(multiply("123", "456"))
}
