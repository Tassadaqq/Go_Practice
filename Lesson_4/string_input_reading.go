package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input your String:\n")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}
	input = strings.TrimSpace(input)
	fmt.Printf("Hello %s!\n", input)

}
