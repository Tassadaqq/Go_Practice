package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randInt := rand.Intn(7)

	for randInt == 0 {
		randInt = rand.Intn(7)
	}

	fmt.Println("Dice rolled, You got\n", randInt)
}
