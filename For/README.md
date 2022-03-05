# GoLang

## For

For é o while do Go.

---

For com condição:
```go 
package main

import (
    "fmt"
)

func main() {
    sum := 0
	for i := 0; i < 10; i++ {
		sum++
		fmt.Println("Soma: ", sum)
	}
}
```

For sem condição (infinito)
```go
package main

import (
    "fmt"
)

func main() {
    value := 0
	for {
		value++
		fmt.Println(value)
	}
}
```