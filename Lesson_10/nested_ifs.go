//Yes it is totally possible

package main

import (
	"fmt"
	"os"
)

func main() {
	var x float32
	fmt.Print("Enter the number you want to enter\n")
	fmt.Fscan(os.Stdin, &x)
	if x > 0 {
		if x > 1000 {
			fmt.Printf("Enter a number less than 1000")
		} else {
			answer := x / 2
			fmt.Printf("Your answer is %f ", answer)
		}
	} else {
		fmt.Printf("Enter a number greater than zero")
	}
}
