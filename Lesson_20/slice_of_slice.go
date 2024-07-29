package main
import "fmt"
func main() {
	mystring := "Hello World!"
	s1 := mystring[0:5]
	s2 := mystring[6:12]
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := s1[1:3]
	fmt.Println(s3)
}
