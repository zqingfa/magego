package main

import "fmt"

func main() {
	var slice = []int{9, 8, 19, 10, 2, 8}
	for i := len(slice) - 1; i >= 0; i-- {
		for j := i; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}
