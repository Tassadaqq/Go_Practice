package main

import (
	"fmt"
)

func main() {
	var arr = [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
