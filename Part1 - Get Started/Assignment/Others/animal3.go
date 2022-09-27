

package main
import (
	"fmt"
)

// Define an interface type called Animal which describes the methods of an animal. 
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), 
// which take no arguments and return no values. 
type Animal interface {
 	Eat() 
	Move()
	Speak()
}

// Define three types Cow, Bird, and Snake. 
// For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird,
//  and Snake all satisfy the Animal interface.
// {"grass", "walk", "moo"}
// {"worms", "fly", "peep"}
// {"mice", "slither", "hsss"}
type Cow struct {
	food string
	locomotion string
	noise string
}
func (a Cow) Eat() {
	fmt.Println(a.food)
}
func (a Cow) Move() {
	fmt.Println(a.locomotion)
}
func (a Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food string
	locomotion string
	noise string
}
func (a Bird) Eat() {
	fmt.Println(a.food)
}
func (a Bird) Move() {
	fmt.Println(a.locomotion)
}
func (a Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food string
	locomotion string
	noise string
}
func (a Snake) Eat() {
	fmt.Println(a.food)
}
func (a Snake) Move() {
	fmt.Println(a.locomotion)
}
func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func main () {
	// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
	// Your program should accept one command at a time from the user, print out a response,
	//  and print out a new prompt on a new line. Your program should continue in this loop forever.
	// Every command from the user must be either a “newanimal” command or a “query” command.
	var sst, snd, srd string
	animals := map[string]int{
		"bird": 1,
		"snake": 2,
		"cow": 3,
	}
	actions := map[string]int{
		"eat": 1,
		"move": 2,
		"speak": 3,
	}

	zoo := make(map[string]Animal)
	for {
		fmt.Printf(">\t")
		fmt.Scanf("%s %s %s", &sst, &snd, &srd)

		// Each “newanimal” command must be a single line containing three strings.
		// The first string is “newanimal”. 
		// The second string is an arbitrary string which will be the name of the new animal.
		// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  
		// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
		if sst == "newanimal" {
			fmt.Println("ok, newanimal. name: ", snd," and kind: ", srd)
			_, exists  := animals[srd]
			// fmt.Println(exists)
			if !exists {
				fmt.Println("unsupported animal")
				continue
			}
			switch srd {
				case "cow":
					zoo[snd]= Cow{"grass", "walk", "moo"}
				case "bird":
					zoo[snd]= Bird{"worms", "fly", "peep"}
				case "snake":
					zoo[snd]= Snake{"mice", "slither", "hsss"}
			}
			fmt.Println("created it!")
			
		// Each “query” command must be a single line containing 3 strings.
		// The first string is “query”. 
		// The second string is the name of the animal. 
		// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
		// Your program should process each query command by printing out the requested data.
		} else if sst == "query" {
			fmt.Println("ok, query. name: ", snd," and kind: ", srd)
			_, exists  := zoo[snd]
			if !exists {
				fmt.Println("unkown animal, wrong unique name?")
				continue
			}
			_, exists2  := actions[srd]
			if !exists2 {
				fmt.Println("unkown action. eat, speak, move.")
				continue
			}
			ani := zoo[snd]
			switch srd {
			case "eat":
				ani.Eat()
			case "move":
				ani.Move()
			case "speak":
				ani.Speak()
			}

		} else {
			fmt.Println("wrong comand")
			continue
		} 
	}
}
		
		// if s_animal == "cow" {
		// 	animal = Animal{"grass", "walk", "moo"}
		// } else if s_animal == "bird" {
		// 	animal = Animal{"worms", "fly", "peep"}
		// } else if s_animal == "snake" {
		// 	animal = Animal{"mice", "slither", "hsss"}
		// } else {
		// 	fmt.Println("unknown animal")
		// }

		// if s_req == "eat" {
		// 	animal.Eat()
		// } else if s_req == "move" {
		// 	animal.Move()
		// } else if s_req == "speak" {
		// 	animal.Speak()
		// } else {
		// 	fmt.Println("unknown request")
		// }