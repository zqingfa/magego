package main

import "fmt"

func main() {
	//插入排序
	arr := []int{10, 2, 3, 4, 5, 7, 5, 8}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
