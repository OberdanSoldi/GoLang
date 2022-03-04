# GoLang

## 2. Pacotes

```go
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    fmt.Println("Meu número é: ", rand.Intn(100)) // Gera o mesmo número por execução
    rand.Seed(time.Now().UnixNano())
	fmt.Println("Número: ", rand.Int()) // Gera um número aleatório por execução
}
```