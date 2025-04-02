package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 || (strings.ToUpper(os.Args[1]) != "C" && strings.ToUpper(os.Args[1]) != "F") {
		printError("Invalid arguments. Usage: program [C|F]")
	}

	originUnit := strings.ToUpper(os.Args[1])
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("What is the current temperature in %s? ", originUnit)
		input, err := reader.ReadString('\n')
		if err != nil {
			printError("Error reading input")
		}

		input = strings.TrimSpace(input)
		originValue, err := strconv.ParseFloat(input, 64)
		if err != nil {
			printError("Invalid temperature value. Please enter a number")
			continue
		}

		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature? (y/n) ")
		continueResponse, err := reader.ReadString('\n')
		if err != nil {
			printError("Error reading input")
		}

		continueResponse = strings.TrimSpace(continueResponse)
		if strings.ToUpper(continueResponse) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%.1f F = %.1f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%.1f C = %.1f F\n", value, convertedValue)
}
