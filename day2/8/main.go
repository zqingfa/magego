package main

import "fmt"

func inputNum(a *int) {
	_, err := fmt.Scan(a)
	if err != nil {
		return
	}

}
func main() {
	var slice1 = []int{2, 8, 9, 10, 19}
	var selectValue int
	fmt.Print("输入查找数字：")
	inputNum(&selectValue)
	fmt.Println(selectValue)

	var middleIndex int
	var startIndex = 0
	var endIndex = len(slice1) - 1
	for startIndex <= endIndex {
		middleIndex = (startIndex + endIndex) / 2
		if selectValue == slice1[middleIndex] {
			fmt.Println("索引是：", middleIndex)
			break
		} else if selectValue > slice1[middleIndex] {
			startIndex = middleIndex + 1
		} else if selectValue < slice1[middleIndex] {
			endIndex = middleIndex - 1
		} else {

		}
	}
	if startIndex > endIndex {
		fmt.Println("数据不存在")
	}

}
