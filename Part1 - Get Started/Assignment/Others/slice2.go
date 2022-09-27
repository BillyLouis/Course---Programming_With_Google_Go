package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	x := make([]int, 0, 3)
	for {
		var t string
		fmt.Printf("Please enter your next integer: ")
		fmt.Scan(&t)

		if strings.Compare(t, "X") == 0 {
			break
		}

		i, err := strconv.ParseInt(t, 0, 64)
		if err != nil {
			fmt.Printf("Wrong Input. \n")
			continue
		}

		x = append(x, int(i))
		sort.Ints(x)

		fmt.Printf("Current Slice is: %v\n", x)

	}
}
