package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Meu número é: ", rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Número: ", rand.Int())
}
