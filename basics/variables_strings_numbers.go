package main 

import "fmt"

func main() {
	// strings
	var nameOne string = "Free"
	var nameTwo = "Palenstine"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "Palenstine must win"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yeaaaaaaaaaaaaa"

	fmt.Println(nameFour)

	// numbers
	var ageOne int = 21
	var ageTwo = 22
	var ageThree int		// default is 0

	fmt.Println(ageOne, ageTwo, ageThree)

	ageFour := 23
	fmt.Println(ageFour)
}