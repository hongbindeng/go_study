package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var pc [256]byte

func pcInit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func SprintTestJon(terms []string) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(q)
}

func main() {
	var valueTest uint64
	valueTest = 1234
	pcInit()
	intTest := PopCount(valueTest)
	fmt.Println(intTest)
	fmt.Println(pc)
	SprintTestJon(os.Args[1:])
	// var s, seq string
	// //var test_arg string
	// silce := make([]int, 3, 5)
	// test_arg := silce[1:]
	// fmt.Println(test_arg)

	// for i := 0; i < len(os.Args); i++ {
	// 	s += seq + os.Args[i]
	// 	seq += " "
	// }
	//fmt.Println(s)
}
