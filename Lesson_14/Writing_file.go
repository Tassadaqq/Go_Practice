package main

import "os"

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("Lahore\n")
	file.WriteString("Okara\n")
	file.WriteString("Karachi\n")
	file.WriteString("Islamabad\n")
}
