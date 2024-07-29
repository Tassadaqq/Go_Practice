package main
import "fmt"
func main() {
	var x int = 5
	var ipointer *int
	ipointer = &x
	fmt.Printf("Memory address of x: %x\n", &x)
	fmt.Printf("Memory address of ipointer: %x\n", ipointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *ipointer)
}
