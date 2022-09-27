// Write a program which prompts the user to first enter a name, and then enter an address. 
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, 
// respectively. Your program should use Marshal() to create a JSON object from the map, and then your program 
// should print the JSON object.

package main

import (
	"encoding/json";
	"fmt";
	"bufio";
	"os"
)

func main() {

	var nameInput string
	var addressInput string

	fmt.Printf("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    nameInput = scanner.Text()

	fmt.Printf("Please enter your adress: ")
    scanner.Scan() 
    addressInput = scanner.Text()
	
	m := make(map[string]string)

	m["name"] = nameInput
	m["address"] = addressInput

	prettyJSON, err := json.MarshalIndent(m, "", "    ")
	
    if err != nil {
        fmt.Println("JSON parse error: ", err)
        return
	}
	
    fmt.Println(string(prettyJSON))
}
