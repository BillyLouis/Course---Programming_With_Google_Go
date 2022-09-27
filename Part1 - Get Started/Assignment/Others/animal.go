package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}


func main() {

	fmt.Println("Enter 'exit' to Quit")
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter request: ")

	for ;; {
		
		fmt.Print("> ")		
		input, _ := reader.ReadString('\n') 

		if (input == "exit") {
			fmt.Println("should exit now")
			break
		}

		var obj Animal
		req_arr := strings.Fields(input)

		if ((len(req_arr) == 1) && (req_arr[0] == "exit")) {
			fmt.Println("should exit now")
			break
		}

		if (len(req_arr) != 2) {
			continue
		}

		if (req_arr[0] == "cow") {
				obj = Animal{"grass", "walk", "moo"}
		}
		if (req_arr[0] == "bird") {
				obj = Animal{"worms", "fly", "peep"}
		}
		if (req_arr[0] == "snake") {
				obj = Animal{"mice", "slither", "hsss"}
		}

		if (req_arr[1] == "eat") {
			obj.Eat()
		}
		if (req_arr[1] == "move") {
			obj.Move()
		}
		if (req_arr[1] == "speak") {
			obj.Speak()
		}
	}
}