package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	inp := bufio.NewScanner(os.Stdin)
	fmt.Println("File name (example.txt):")
	inp.Scan()
	filename := inp.Text()
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	people := make([]Person, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " ")
		people = append(people, Person{splitted[0], splitted[1]})
	}
	for _, person := range people {
		fmt.Printf("%s\n%s\n\n", person.fname, person.lname)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
