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

	// numbers and integers
	var ageOne int = 21
	var ageTwo = 22
	var ageThree int		// default is 0

	var age int8			// default is 0
	fmt.Println("age: ", age)

	fmt.Println(ageOne, ageTwo, ageThree)

	ageFour := 23
	fmt.Println(ageFour)

	// var numOne int8 = 1327		// cannot use 1327 (untyped int constant) as int8 value in variable declaration (overflows)
	// fmt.Println(numOne)

	var numOne int8 = 125		// -128 to 127
	fmt.Println("numOne: ", numOne)

	var numTwo uint8 = 241		// 0 to 255
	fmt.Println("numTwo: ", numTwo)

	var numThree int16 = 664		// -32768 to 32767
	fmt.Println("numThree: ", numThree)

	var numFour uint16 = 5437		// 0 to 65535
	fmt.Println("numFour: ", numFour)

	var numFive int32 = 444457		// -2147483648 to 2147483647
	fmt.Println("numFive: ", numFive)

	var numSix uint32 = 557		// 0 to 4294967295
	fmt.Println("numSix: ", numSix)

	var numSeven int64 = 1145654		// -9223372036854775808 to 9223372036854775807
	fmt.Println("numSeven: ", numSeven)

	var numEight uint64 = 125		// 0 to 18446744073709551615
	fmt.Println("numEight: ", numEight)

	numNine := 128						// -9223372036854775808 to 9223372036854775807
	fmt.Println("numNine: ", numNine)

	numTen := 5646354354546            	// -9223372036854775808 to 9223372036854775807
	fmt.Println("numTen: ", numTen)

	numEleven := -7646354354546        	// -9223372036854775808 to 9223372036854775807
	fmt.Println("numEleven: ", numEleven)

	// floating point numbers
	var floatOne float32 = 3.14
	var floatTwo = 3.141592

	fmt.Println("floatOne: ", floatOne, "floatTwo: ", floatTwo)

	var floatThree float64 = 3141592.653545
	fmt.Println("floatThree: ", floatThree)

	floatFour := 31415.92653545
	fmt.Println("floatFour: ", floatFour)

	// complex numbers
	var complexOne complex64 = 3 + 2i		// complex64 is the set of all complex numbers with float32 real and imaginary parts.
	fmt.Println("complexOne: ", complexOne)

	var complexTwo complex128 = 3 + 2i		// complex128 is the set of all complex numbers with float64 real and imaginary parts.
	fmt.Println("complexTwo: ", complexTwo)

	complexThree := 3 + 2i
	fmt.Println("complexThree: ", complexThree)

	// boolean
	var boolOne bool = true
	fmt.Println("boolOne: ", boolOne)

	var boolTwo = false
	fmt.Println("boolTwo: ", boolTwo)

	boolThree := true
	fmt.Println("boolThree: ", boolThree)

	// Byte
	var byteOne byte = 'A'		// byte is an alias for uint8
	fmt.Println("byteOne: ", byteOne) // ascii table for 'A' is 65

	var byteTwo = 'B'
	fmt.Println("byteTwo: ", byteTwo)	// ascii table for 'B' is 66

	byteThree := 'C'
	fmt.Println("byteThree: ", byteThree)	// ascii table for 'C' is 67

	// rune
	var runeOne rune = 'a'		// rune is an alias for int32
	fmt.Println("runeOne: ", runeOne) // ascii table for 'a' is 97

	var runeTwo = 'b'
	fmt.Println("runeTwo: ", runeTwo) // ascii table for 'b' is 98

	runeThree := 'c'
	fmt.Println("runeThree: ", runeThree) // ascii table for 'c' is 99



}