package main
import "os"
func main() {
	src := "file.txt"
	dst := "go.txt"
	os.Rename(src, dst)
}
