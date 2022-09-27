package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	text, _ := reader.ReadString('\n')
	t := strings.ToLower(text)
	if strings.HasPrefix(t, "i") && strings.HasSuffix(t, "n\n") && strings.Contains(text, "a") {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")

}

func main() {
	findian()
}
