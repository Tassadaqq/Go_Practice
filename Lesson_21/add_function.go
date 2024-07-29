package main

import "fmt"
func main() {
	a := 4
	b := 8
	result := add(a, b)
	fmt.Println(result)
}
func add(a int, b int) int {
	return a + b
}
