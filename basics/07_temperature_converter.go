package basics

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/fatih/color"

)

func TemperatureConverter() {
	fmt.Println("Temperature Converter")
	fmt.Println("Enter the temperature in Celsius and I will convert it to Fahrenheit and Kelvin")
	for {
		fmt.Println("Enter the temperature in Celsius (or 'quit' to exit): ")
		reader := bufio.NewReader(os.Stdin)
		celsius, _ := reader.ReadString('\n')
		celsius = strings.TrimSpace(celsius)
		
		if celsius == "quit" || celsius == "exit" {
			color.Blue("Goodbye!")
			break
		}
		
		celsiusFloat, err := strconv.ParseFloat(celsius, 64)
		if err != nil {
			color.Red("Invalid input for temperature. Please enter a valid number. \n")
			continue
		}
		
		kelvin := celsiusFloat + 273.15
		fahrenheit := (celsiusFloat * 9 / 5) + 32
        color.Green("Celsius: %g, Fahrenheit: %g, Kelvin: %g", celsiusFloat, fahrenheit, kelvin)
	}
}