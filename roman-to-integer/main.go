package main

import (
	"fmt"
)

var base = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
var mapping = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var result int
	var length = len(s)
	for i := 0; i < length; {
		if i+2 <= length {
			token := s[i : i+2]
			val, ok := mapping[token]
			if ok { //match!
				result += val
				i += 2
				continue
			}
		}
		result += base[s[i]]
		i++

	}
	return result
}
