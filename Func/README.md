# GoLang

## Func

```go
package main

import (
    "fmt"
)

func add(x int, y int) int { //Tipo da variável declarada depois do nome
    x + y
}

func main() {
    fmt.Println("Soma 1 + 1 =", add(1, 1))
	fmt.Println("Soma 3 + 3 =", add(3, 3))
	fmt.Println("Soma 4 + 4 =", add(4, 4))
}
```