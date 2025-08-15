package basics

import "fmt"

func ArraysAndSlices() {

	// Arrays
	var ages [3]int = [3]int{20, 25, 30}

	fmt.Println(ages, len(ages), cap(ages))

	var names1 = [3]string{"John", "Jane", "Joe"}
	names1[1] = "Jim"
	fmt.Println("names1", names1, len(names1))

	ages2 := [3]int{50, 55, 60}
	fmt.Println(ages2, len(ages2))

	// Slices (dynamic arrays)
	scores1 := []float32{45, 55.34, 65.23}
	fmt.Println(scores1, len(scores1))
	scores1[2] = 60.23
	scores1 = append(scores1, 70.23)
	fmt.Println("score1", scores1, len(scores1))

	// Slicing ranges
	scores2 := scores1[1:3] // [low:high]; high is not included
	fmt.Println("slicing ranges", scores2, len(scores2))

	// Slicing with step
	scores3 := scores1[1:3:4] // [low:high:max]; high is not included
	fmt.Println("score3", scores3, len(scores3))

	rangeOne := names1[1:3]
	rangeTwo := names1[2:]
	rangeThree := names1[:2]
	fmt.Println("rangeOne: ", rangeOne, "rangeTwo: ", rangeTwo, "rangeThree: ", rangeThree)

	rangeOne = append(rangeOne, "Jill")
	fmt.Println(rangeOne)

}
