//yes we can call an external function in a recursive function as you can see below
package main
import "fmt"
func calc(x uint) uint {
	return x * factorial(x-1)
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	y := calc(x)
	return y
}
func main() {
	x := factorial(3)
	fmt.Println(x)
}
