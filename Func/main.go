package main

import (
	"fmt"
)

func add(x int, y int) int { // Tipo sempre dopois do nome da variável
	return x + y
}

func main() {
	fmt.Println("Soma 1 + 1 =", add(1, 1))
	fmt.Println("Soma 3 + 3 =", add(3, 3))
	fmt.Println("Soma 4 + 4 =", add(4, 4))
}
