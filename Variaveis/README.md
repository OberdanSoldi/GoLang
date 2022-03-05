# GoLang

## Variáveis

Variáveis com inicializadores
```go
package main

import (
    "fmt"
)

//se o inicializador estiver presente o tipo podera ser omitido, a variável tera o tipo do inicializador

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no"
	fmt.Println(i, j, c, python, java)
}
```
Declaração curta de variáveis
```go
package main

import (
    "fmt"
)

func main() {
	var i, j int = 1, 2
	k := 3 //declaração curta
	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)
}
```