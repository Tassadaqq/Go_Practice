package main

import (
	"fmt"
	"os"
)

func main() {
	var date, year, age int
	fmt.Print("Enter the date you were born on: \n")
	fmt.Fscan(os.Stdin, &date)
	fmt.Print("Enter your age: \n")
	fmt.Fscan(os.Stdin, &age)
	if (age <= 0) || (date >= 32 && date <= 0) {
		fmt.Print("Kindly enter valid date and age\n")
	} else {
		year = 2024 - age
		fmt.Printf("Your year of birth is : %d !", year)
	}

}
