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
	fmt.Println("zuida:", maxNum)
	fmt.Println("suoyinwei :", maxNumIndex)
}
