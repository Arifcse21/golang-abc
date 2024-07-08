package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, GoLang!")
	fmt.Println("I am ", os.Getenv("USER"))
}
