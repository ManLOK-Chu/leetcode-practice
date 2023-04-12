package main

import "fmt"

func StairsCharge(num int) (total float64) {
	var stairs = []int{0, 5, 20, 50, 100, 500, 1000, 2000, 3000, 4000, 5000, 6000}
	var unitPrice = []float64{0, 30, 15, 10, 8, 7, 6, 5, 4, 3, 2, 1}

	var cached = make([]float64, 11)
	for i := 1; i < len(cached); i++ {
		//预处理
		cached[i] = float64(stairs[i]-stairs[i-1])*unitPrice[i] + cached[i-1]
	}
	//fmt.Println(cached)
	for i := 1; i < len(stairs); i++ {
		if stairs[i] >= num {
			total = cached[i-1] + float64(num-stairs[i-1])*unitPrice[i]
			return
		}
	}
	total = cached[len(cached)-1] + float64(num-stairs[len(stairs)-1])*unitPrice[len(unitPrice)-1]
	return
}

func main() {
	fmt.Println(StairsCharge(0))
	fmt.Println(StairsCharge(6))
	fmt.Println(StairsCharge(20))
	fmt.Println(StairsCharge(6001))
}
