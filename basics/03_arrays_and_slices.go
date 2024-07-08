package main

import "fmt"

func main() {

	// Arrays
	var ages [3]int = [3]int{20, 25, 30}

	fmt.Println(ages, len(ages), cap(ages)) 	// cap = capacity

	var names1 = [3]string{"John", "Jane", "Joe"}
	names1[1] = "Jim"
	fmt.Println(names1, len(names1), cap(names1)) 	// cap = capacity

	ages2 := [3]int{50, 55, 60}
	fmt.Println(ages2, len(ages2), cap(ages2)) 	// cap = capacity


	// Slices (dynamic arrays)
	scores1 := []float32{45, 55.34, 65.23}
	fmt.Println(scores1, len(scores1), cap(scores1)) 	// cap = capacity
	scores1[2] = 60.23
	scores1 = append(scores1, 70.23)
	fmt.Println(scores1, len(scores1), cap(scores1)) 	// cap = capacity

	// Slicing ranges
	scores2 := scores1[1:3]
	fmt.Println(scores2, len(scores2), cap(scores2)) 	// cap = capacity

	// Slicing with step
	scores3 := scores1[1:3:4]	// [begin:end:step]
	fmt.Println(scores3, len(scores3), cap(scores3)) 	// cap = capacity

	rangeOne := names1[1:3]
	rangeTwo := names1[2:]
	rangeThree := names1[:2]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Jill")
	fmt.Println(rangeOne)

}