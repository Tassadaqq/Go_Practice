//yes a combination of different returns can be used
package main
import "fmt"
func values() (int, string) {
	return 5, "world!"
}
func main() {
	x, y := values()
	fmt.Println(x)
	fmt.Println(y)
}
