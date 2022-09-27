package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create struct
type Name struct {
	fname string
	lname string
}

func main() {

	fmt.Println("Please, write the name of the file.")

	// Scan file name
	var filename string
	fmt.Scan(&filename)

	// Create empty slice to store names
	var nameSlice []Name

	// Open file and return error if not found
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)

	}

	// Close the file in the end of execution
	defer f.Close()

	// Initialize scan for each line
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Variables for loop
	var nameStruct Name
	var line []string
	// Loop for lines and store info in slice
	for scanner.Scan() {

		// Split line by spaces and save it in slice
		line = strings.Split(scanner.Text(), " ")

		// Fill struct with 2 first elements
		nameStruct.fname = line[0]
		nameStruct.lname = line[1]

		// Append slice to struct
		nameSlice = append(nameSlice, nameStruct)

	}

	// Loop for struct and print values
	for _, val := range nameSlice {
		fmt.Println(val)
	}

}
