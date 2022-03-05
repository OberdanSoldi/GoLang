package main

import (
	"fmt"
)

func add(x int, y int) int { // ou pode ser declarado como (x, y int)
	// Tipo sempre dopois do nome da variável
	return x + y
}

func swap(x, y string) (string, string) { //A função pode retornar qualquer número de dados
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 1
	y = sum - x
	return
}

func main() {
	fmt.Println("Soma 1 + 1 =", add(1, 1))
	fmt.Println("Soma 3 + 3 =", add(3, 3))
	fmt.Println("Soma 4 + 4 =", add(4, 4))
	//-------------------------------------//
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//-------------------------------------//
	fmt.Println(split(17))
}
