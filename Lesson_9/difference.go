package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30}
	for index, element := range a {
		fmt.Printf("Index: %d, Element: %d\n", index, element)
	}

}
// in this we are including both index and element to the range of array a this will give us both the index at the moment and the value which is inside of that array

package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30}
	for _,element := range a {
		fmt.Printf(Element: %d\n", element)
	}

}
\\ the _, in go language is used to filter out the elements which we do not want in above case we will only get the index values rather than the index number as we do not want it
