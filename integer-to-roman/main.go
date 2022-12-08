package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	b := strings.Builder{}
	for num != 0 {
		switch {
		case num >= 1 && num < 4:
			b.WriteByte('I')
			num -= 1
		case num == 4:
			b.WriteString("IV")
			num -= 4
		case num >= 5 && num < 9:
			b.WriteByte('V')
			num -= 5
		case num == 9:
			b.WriteString("IX")
			num -= 9
		case num >= 10 && num < 40:
			b.WriteByte('X')
			num -= 10
		case num >= 40 && num < 50:
			b.WriteString("XL")
			num -= 40
		case num >= 50 && num < 90:
			b.WriteByte('L')
			num -= 50
		case num >= 90 && num < 100:
			b.WriteString("XC")
			num -= 90
		case num >= 100 && num < 400:
			b.WriteByte('C')
			num -= 100
		case num >= 400 && num < 500:
			b.WriteString("CD")
			num -= 400
		case num >= 500 && num < 900:
			b.WriteByte('D')
			num -= 500
		case num >= 900 && num < 1000:
			b.WriteString("CM")
			num -= 900
		case num >= 1000:
			b.WriteByte('M')
			num -= 1000
		}
	}
	return b.String()
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(500))
	fmt.Println(intToRoman(1994))
}
