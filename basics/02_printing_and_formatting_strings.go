package main

import (
	"fmt"
	"os"
)

func main() {
	age := 20
	name := os.Getenv("USER")

	// Print, no new line
	fmt.Print("Hello!, ")
	fmt.Print("I am ")
	fmt.Print(name)
	fmt.Print(". and I am ")
	fmt.Print(age)
	fmt.Print(" years old.")
	fmt.Print("\n\n")
	
	// Print, with new line
	fmt.Println("Hello!, ")
	fmt.Println("I am", name, ".")		// comma separated adds a space, mind it, sweetie!
	fmt.Println("and I am", age, "years old.")

	// Printf  (formatted string) %_ = format specifier
	fmt.Printf("Hello!, I am %v. and I am %v years old.\n", name, age)				// %v = value, default for variables
	fmt.Printf("Hello!, I am %s. and I am %d years old.\n", name, age)				// %s = string, %d = int
	fmt.Printf("Hello!, I am %q. and I am %q years old.\n", name, age)				// %q = quoted string, doesn't work with numbers. Useful for strings with spaces.
	fmt.Printf("name is a type of %T. and age is a type of %T .\n", name, age)		// %T = type
	fmt.Printf("You scored %f points.\n", 10.5652)									// %f = float
	fmt.Printf("You scored %F points.\n", 10.5652)									// %F = synonym for %f
	fmt.Printf("You scored %0.2f points.\n", 10.5652)								// %0.2f = float with 2 decimals

	fmt.Printf("Hello!, I am %p. and I am %p years old.\n", name, age)				// %p = address of 0th element in base 16 notation, with leading 0x

	

	// Sprintf
	s := fmt.Sprintf("Hello!, I am %s. and I am %d years old.", name, age)
	fmt.Println(s)

	// Fprintf
	fmt.Fprintf(os.Stderr, "Hello!, I am %s. and I am %d years old.\n", name, age)



}