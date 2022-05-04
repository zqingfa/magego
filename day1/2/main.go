package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	a, b = 1, 1
	for a = 1; a <= 9; a++ {
		for b = 1; b <= a; b++ {
			fmt.Printf("%d*%d=%d\t", a, b, a*b)
		}
		fmt.Println()
	}

}
