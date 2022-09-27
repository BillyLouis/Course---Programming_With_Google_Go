package main

// Seach one action about the animal and identify which animal

import (
	"fmt"
	"strings"
)

type Animal struct{ food, locomotion, noise string }

func (e *Animal) Eat() {
	fmt.Print(e.food)
}

func (m *Animal) Move() {
	fmt.Print(m.locomotion)
}

func (s *Animal) Speak() {
	fmt.Print(s.noise)
}

func validateAction(action string, animal Animal) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "Speak":
		animal.Speak()
	default:
		println("No valid action")
	}
}

func main() {
	var theAnimal string
	var search string

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print("\nSearch  e.g cow  move > ")
		fmt.Scan(&theAnimal, &search)

		theAnimal = strings.ToLower(theAnimal)
		search = strings.ToLower(search)

		switch theAnimal {
		case "cow":
			fmt.Print(theAnimal, " ")
			validateAction(search, cow)
		case "bird":
			fmt.Print(theAnimal, " ")
			validateAction(search, bird)
		case "snake":
			fmt.Print(theAnimal, " ")
			validateAction(search, snake)
		default:
			println("ERROR -> Not a valid animal")
		}

	}

}
