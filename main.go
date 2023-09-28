package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(f)
	}
}
