package main

import "fmt"

func main() {
	var slience3 = []int{9, 8, 19, 10, 2, 8}
	var maxNum int
	var secondNumIndex int
	var secondNum int
	for k, v := range slience3 {
		if v > maxNum {
			maxNum = v

		} else if v > secondNum {
			secondNum = v
			secondNumIndex = k
		}
	}
	for i := secondNumIndex; i < len(slience3)-2; i++ {
		slience3[i] = slience3[i+1]
	}
	slience3[len(slience3)-2] = secondNum
	fmt.Println("输出后为:", slience3)
	//
	//	fmt.Println("第二大:", secondNum)
	//	fmt.Println("第二大索引 :", secondNumIndex)
}
