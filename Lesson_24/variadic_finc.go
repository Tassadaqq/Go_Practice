package main
import "fmt"
func variadic_func(students ...string) {
	for _, student := range students {
		fmt.Println(student)
	}
}
func main() {
	variadic_func("Tassadaq", "Abid", "Basit", "Tariq", "Khalid")
}
