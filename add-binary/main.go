package main

import "fmt"

func addBinary(a string, b string) string {
	var short, long = b, a
	var lenShort, lenLong = len(b), len(a)
	var result = make([]byte, lenLong+1)
	if lenShort > lenLong {
		short, long = a, b
		lenShort, lenLong = lenLong, lenShort
	}
	var tmp = make([]byte, lenLong)
	for i := 0; i < (lenLong - lenShort); i++ {
		tmp[i] = '0'
	}
	tmp = append(tmp, []byte(short)...)
	for i := range tmp {
		result[i] = tmp[i] - '0' + long[i] - '0'
	}
	for i := range result {
		result[i] += '0'
	}
	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	//fmt.Println(addBinary("1010", "1011"))
	//fmt.Println(addBinary("0", "0"))
	//fmt.Println(addBinary("1", "1"))
	//fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
