package basics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Devide(a, b float64) float64 {
	if b == 0 {
		panic("Cannot devide by zero")
	}
	return a / b
}

func PanicFunction() {
	fmt.Println("Panic Function")
	fmt.Println("Enter first number: ")
	reader1 := bufio.NewReader(os.Stdin)
	num1Str, _ := reader1.ReadString('\n')
	num1, _ := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)
	fmt.Println("You entered: ", num1)

	fmt.Println("Enter second number: ")
	reader2 := bufio.NewReader(os.Stdin)
	num2Str, _ := reader2.ReadString('\n')
	num2, _ := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)
	fmt.Println("You entered: ", num2)
	
	result := Devide(num1, num2)
	fmt.Println("Result: ", result)
}
