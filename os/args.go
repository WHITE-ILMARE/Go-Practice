package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Args count = %d\n", len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
