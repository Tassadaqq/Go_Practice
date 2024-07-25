package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter a number between 1 and 10 \n")
	fmt.Scanf("%d", &age)
	if age >= 10 || age <= 1 {
		fmt.Printf("Kindly enter number between 1 and 10\n")
	} else {
		fmt.Printf("The number you entered is within the range Thankyou\n")
	}

}
