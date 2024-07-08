package main

import  (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you city: ")
	city, _ := reader.ReadString('\n')
	fmt.Println("Your city is: ", city)

	
}