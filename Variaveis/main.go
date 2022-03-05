package main

import (
	"fmt"
)

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

//-------------------------------------------//

// Variáveis com inicializadores

// se o inicializador estiver presente o tipo podera ser omitido, a variável tera o tipo do inicializador

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no"
// 	fmt.Println(i, j, c, python, java)
// }

//-------------------------------------------//

// Declarações curtas de variáveis

// func main() {
// 	var i, j int = 1, 2
// 	k := 3 //declaração curta
// 	c, python, java := true, false, "no"

// 	fmt.Println(i, j, k, c, python, java)
// }
