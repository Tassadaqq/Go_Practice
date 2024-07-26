package main

import (
	"fmt"
	"os"
)

func main() {

		if _, err := os.Stat("D:/program.go"); err == nil {
			fmt.Printf("File exists\n")
		} else {
			fmt.Printf("File does not exist\n")
		}
	}
