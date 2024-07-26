package main

import (
	"fmt"
)

type House struct {
	noRooms int
	price   float32
	city    string
}

func print(h House) {
	fmt.Printf("The house price is %f the city is %s and number of rooms is %d!\n", h.price, h.city, h.noRooms)
}

func main() {
	var h1 House
	h1.city = "Lahore"
	h1.noRooms = 25
	h1.price = 125000000
	print(h1)

	var h2 House
	h2.city = "New York"
	h2.noRooms = 4
	h2.price = 54000000
	print(h2)
}
