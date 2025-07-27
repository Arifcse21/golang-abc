package basics

import  (
	"bufio"
	"fmt"
	"os"
)

func KeyboardInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you city: ")
	city, _ := reader.ReadString('\n')
	fmt.Println("Your city is: ", city)

	
}