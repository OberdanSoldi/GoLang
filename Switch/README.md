# GoLang

## Switch

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on: ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}
```