package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	wordCount := make(map[byte]int)
	sourceFile, err := ioutil.ReadFile("IHaveADream.txt")
	if err != nil {
		fmt.Println("open file failed")
	}

	for _, v := range sourceFile {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			wordCount[v]++
		}
	}
	/*
	   排序输出
	*/
	/*
			for i := 'a'; i <= 'z'; i++ {
			fmt.Printf("%c:次数为%v \n", i, wordCount[byte(i)])
		}
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c:次数为%v \n", i, wordCount[byte(i)])
		}
	*/

	/*
	   乱序输出
	*/
	for k, v := range wordCount {
		fmt.Printf("%c:次数为%v \n", k, v)
	}
}
