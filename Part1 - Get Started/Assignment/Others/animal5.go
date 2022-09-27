package main

import "fmt"

type Animal interface {
	Speak()
	Move()
	Eat()
}

type Cow struct {
	food, sound, locomotion, name string
}

func (a Cow) Speak() {
	fmt.Printf("%s says %s\n", a.name, a.sound)
}
func (a Cow) Move() {
	fmt.Printf("%s can %s\n", a.name, a.locomotion)
}
func (a Cow) Eat() {
	fmt.Printf("%s eats %s\n", a.name, a.food)
}

type Bird struct {
	food, sound, locomotion, name string
}

func (a Bird) Speak() {
	fmt.Printf("%s says %s\n", a.name, a.sound)
}
func (a Bird) Move() {
	fmt.Printf("%s can %s\n", a.name, a.locomotion)
}
func (a Bird) Eat() {
	fmt.Printf("%s eats %s\n", a.name, a.food)
}

type Snake struct {
	food, sound, locomotion, name string
}

func (a Snake) Speak() {
	fmt.Printf("%s says %s\n", a.name, a.sound)
}
func (a Snake) Move() {
	fmt.Printf("%s can %s\n", a.name, a.locomotion)
}
func (a Snake) Eat() {
	fmt.Printf("%s eats %s\n", a.name, a.food)
}

func findAnimal(name string, animals map[string]Animal) (Animal, bool) {
	a, ok := animals[name]
	return a, ok
}

func performAction(a Animal, action string) {
	switch action {

	case "speak":
		a.Speak()

	case "eat":
		a.Eat()

	case "move":
		a.Move()

	default:
		fmt.Println("Invalid action")
	}
}

func createAnimal(name, species string) (Animal, bool) {
	var output Animal
	ok := false

	switch species {
	case "cow":
		output = Cow{food: "grass", locomotion: "walk", sound: "moo", name: name}
		ok = true
	case "bird":
		output = Bird{food: "worms", locomotion: "fly", sound: "peep", name: name}
		ok = true

	case "snake":
		output = Snake{food: "mice", locomotion: "slither", sound: "hsss", name: name}
		ok = true

	default:
		fmt.Println("Invalid species")
	}

	return output, ok
}

func main() {
	var animals = make(map[string]Animal)

	for {
		var x, y, q string
		fmt.Print("> ")
		fmt.Scanln(&q, &x, &y)

		if q == "newanimal" {
			a, ok := createAnimal(x, y)

			if !ok {
				continue
			}

			fmt.Println("Created it!")
			animals[x] = a
		} else if q == "query" {
			a, ok := findAnimal(x, animals)

			if !ok {
				fmt.Println("Animal not found, try again.")
				continue
			}

			performAction(a, y)

		} else {
			fmt.Println("Invalid query string")
			continue
		}

	}
}
