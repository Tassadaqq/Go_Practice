
package main

import (
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat("C:/Program Files/Go/src/bufio/bufio.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}
