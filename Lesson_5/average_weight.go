package main

import (
	"fmt"
	"os"
)

func main() {
	n := 0
	totalweight := 0
	weight := 0
	fmt.Print("Enter the number of persons you want to add the weight of: \n")
	fmt.Fscan(os.Stdin, &n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Enter the weight of person%d : \n", i)
		fmt.Fscan(os.Stdin, &weight)
		if weight > 0 {
			totalweight += weight
		} else {
			fmt.Print("Enter the weight again\n")
			i--
		}
	}
	average := totalweight / n
	fmt.Printf("The average weight is %d!", average)

}
