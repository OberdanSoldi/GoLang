package main

import (
	"fmt"
)

// For é o while do GO
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
		fmt.Println("Soma: ", sum)
	}

	//caso a condição for omitida vira um laço infinito
	// value := 0
	// for {
	// 	value++
	// 	fmt.Println(value)
	// }
}
