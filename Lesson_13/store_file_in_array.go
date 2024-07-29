package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	file, err := os.Open("name.txt")
	if err != nil {
		return
	}
	var name []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name = append(name, scanner.Text())
	}
	for _, n := range name {
		fmt.Println(n)
	}
}
