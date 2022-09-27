package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers = make([]int, 3)

	var input string

	count := 0

	for {
		// Repeat until X in inputted
		fmt.Println("Enter an integer number or X to quit: ")
		fmt.Scanln(&input)

		// Exit condition
		if string(input) == "X" {
			break
		}

		// Convert the value
		converted_int, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("INVALID INPUT ERROR, Try again...")
			continue
		}

		// Add the number in numbers
		if count < 3 {
			numbers[0] = converted_int
			count = count + 1
		} else {
			numbers = append(numbers, converted_int)
		}

		sort.Ints(numbers)

		fmt.Printf("%v", numbers)
	}

}
