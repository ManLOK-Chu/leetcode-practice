package main

import "fmt"

func compareVersion(version1 string, version2 string) int {
	length1, length2 := len(version1), len(version2)
	var i, j int
	for i < length1 || j < length2 {
		var x, y = 0, 0

		for ; i < length1 && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++ // 跳过点号

		for ; j < length2 && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++ // 跳过点号

		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}
