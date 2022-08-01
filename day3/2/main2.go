package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	wordCount := make(map[byte]int)
	//topWordCount := make(map[int]byte)
	topWordCount := make([]byte, 10)
	sourceFile, err := ioutil.ReadFile("IHaveADream.txt")
	if err != nil {
		fmt.Println("open file failed")
	}
	//将读取的字节转化为大写
	sourceFileUpper := bytes.ToUpper(sourceFile)
	//统计A-Z之间的各个字母频次
	for _, v := range sourceFileUpper {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			wordCount[v]++
		}
	}
	/*
	   排序输出
	*/

	//for i := 'a'; i <= 'z'; i++ {
	//	fmt.Printf("%c:次数为%v \n", i, wordCount[byte(i)])
	//}
	for i := 'A'; i <= 'Z'; i++ {
		for j := 0; j < 9; j++ {
			if wordCount[byte(i)] >= wordCount[topWordCount[j]] {

				//切片的copy方法可以将一个切片中的元素拷贝到另一个切片中
				//如果 src 比 dst 长，就截断；如果 src 比 dst 短，则只拷贝 src 那部分
				copy(topWordCount[j+1:], topWordCount[j:])
				topWordCount[j] = byte(i)
				//后续已无意义，直接退出大小值循环
				break
			}
		}

	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%c:次数为%v \n", topWordCount[i], wordCount[topWordCount[i]])
	}
}
