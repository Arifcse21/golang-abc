package basics

import (
	"fmt"
	"os"
)

func HelloWorld() {
	fmt.Println("Hello, GoLang!")
	fmt.Println("I am ", os.Getenv("USER"))
}
