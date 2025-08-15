package basics

import "fmt"

func ForLoop() {
	// 1. Print Numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println("Number is: ", i)
	}

	// 2. Print Even Numbers 1 to 20
	for i := 2; i <= 20; i += 2 {
		fmt.Println("Even number is: ", i)
	}

	// 3. Print Numbers in Reverse (10 to 1)
	for i := 10; i >= 1; i-- {
		fmt.Println("Reverse number is: ", i)
	}

	// 4. Print Sum of 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100 is: ", sum)

	// 5. Print Multiplication Table of 5
	for i := 1; i <= 10; i++ {
		fmt.Println("5 *", i, "=", 5*i)
	}

	// 6. Iterate over a slice
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums { // for i, v := range []int{1, 2, 3, 4, 5}
		fmt.Println("index:", i, "value:", v)
	}

	// 7. Print Characters in a String
	word := "Golang"
	for i := 0; i < len(word); i++ {
		fmt.Printf("Character: %c \n", word[i])
	}

	// 8. Count Digits of a Number
	num := 870563
	count := 0

	for num > 0 {
		num /= 10
		count++
	}

	fmt.Println("Count of digits is: ", count)

	// 9. Factorial of a Number
	fact := 1
	number := 4
	for i := 1; i <= number; i++ {
		fact *= i
	}

	fmt.Println("Factorial of", number, "is: ", fact)

	// 10. Nested Loop: Print Pattern (Right-Angle Triangle)
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
