package basics

import "fmt"

func DeferFunction() {
	fmt.Println("Defer Function")
	defer fmt.Println("First")
	defer fmt.Println("Second")
	// LIFO Order: If multiple defer statements are present within a single function, 
	// the deferred functions are executed in Last-In, First-Out (LIFO) order. 
	// The last defer statement encountered will be executed first when the surrounding function returns.
	fmt.Println("Third")
	fmt.Println("Fourth")
	defer fmt.Println("Fifth")
	fmt.Println("Sixth") 
}