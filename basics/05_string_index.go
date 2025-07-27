package basics

import (
	"fmt"
	"strings"
)

func StringIndex() {
	s := "Go string example"

	for k, v := range s {
		fmt.Println("k:", k, "v:", v)
	}
	fmt.Println("")
	// Print only even index
	for k := 0; k < len(s); k = k + 2 {
		fmt.Println("k:", k, "v:", s[k])
	}

	// string join
	s1 := "Hello"
	s2 := "Go Lang"
	s3 := s1 + " " + s2

	fmt.Println("\n", s3)

	new_s := strings.Join([]string{s1, s2}, " ")
	fmt.Println(new_s)

}
