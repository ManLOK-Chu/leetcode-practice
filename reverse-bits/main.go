package main

import (
	"fmt"
)

const m = 1<<32 - 1
const m0 = 0x5555555555555555 // 01010101 ...
const m1 = 0x3333333333333333 // 00110011 ...
const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
const m3 = 0x00ff00ff00ff00ff // etc.

func reverseBits(num uint32) uint32 {
	//fmt.Printf("%032b\n", num)
	num = num>>1&(m0&m) | num&(m0&m)<<1
	num = num>>2&(m1&m) | num&(m1&m)<<2
	num = num>>4&(m2&m) | num&(m2&m)<<4
	num = num>>8&(m3&m) | num&(m3&m)<<8
	return num>>16 | num<<16
}

func main() {
	fmt.Printf("%032b\n%032b\n", uint32(43261596), reverseBits(43261596))
	fmt.Printf("%032b\n%032b\n", uint32(4294967293), reverseBits(4294967293))
}
