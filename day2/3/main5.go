package main

import "fmt"

func main() {
	var slience3 = []int{9, 8, 19, 10, 2, 8}
	var maxNum int
	var maxNumIndex int
	for k, v := range slience3 {
		if v > maxNum {
			maxNum = v
			maxNumIndex = k
		}
	}
	for i := maxNumIndex; i < len(slience3)-1; i++ {
		slience3[i] = slience3[i+1]
	}
	slience3[len(slience3)-1] = maxNum
	//slience3Hou := slience3[maxNumIndex+1 : len(slience3)-2]
	fmt.Println("输出后为:", slience3)

}
