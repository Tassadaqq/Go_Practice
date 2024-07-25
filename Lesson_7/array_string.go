package main

import (
	"fmt"
)

func main() {
	var arr = [5]string{"muhammad", "tassadaq", "abid", "is", "great"}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s ", arr[i])
	}
}
