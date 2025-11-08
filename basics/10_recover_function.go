package basics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SafeDevide(a, b float64) (float64, bool) { // Return value and success flag
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Division error recovered:", err)
		}
	}() // expression in defer must be function call

	if b == 0 {
		panic("Cannot divide by zero")
		// return 0, false // This line is unreachable due to panic
	}
	return a / b, true
}

func RecoverFunction() {
	fmt.Println("Recover Function")
	fmt.Println("Enter two numbers and I will divide them")
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
	
	result, success := SafeDevide(num1, num2)
	if success {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("No result due to division by zero")
	}
}