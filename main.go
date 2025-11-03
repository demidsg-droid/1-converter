package main

import "fmt"

func main() {
	const USDtoEUR = 0.87
	const USDtoRUB = 78.84
	const EURtoRUB = (1.0 / USDtoEUR) * USDtoRUB

	fmt.Printf("Стоимость EUR относительно RUB: %.2f", EURtoRUB)

}
