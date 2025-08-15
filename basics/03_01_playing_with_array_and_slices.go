package basics

import "fmt"

func PlayingWithArrayAndSlices() {
	fmt.Println("Playing with array and slices")
	arr1 := []string{"Bal", "Chal", "Khal", "Mal", "Shal", "Thal"}
	fmt.Println(arr1)

	to_mal := arr1[:4]
	fmt.Println(to_mal)
	no_bal := arr1[1:]
	fmt.Println(no_bal)
	only_mal := arr1[3]
	fmt.Println(only_mal)
	khal_to_shal := arr1[2:5]
	fmt.Println(khal_to_shal)

	arr2 := []float32{1.1, 2.234325, 3.323525, 4.235234, 5.2355}
	fmt.Println(arr2)
	snd_float := arr2[1]
	fmt.Printf("the second float is: %0.2f \n", snd_float)
	
	arr3 := append(arr2, 6.2345)
	fmt.Println(arr3)
}
