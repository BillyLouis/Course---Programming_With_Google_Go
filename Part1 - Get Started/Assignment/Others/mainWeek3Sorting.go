package main

// (c) 2022 CAR, All rights reserved

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

func sorter(myID int, elements []int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	fmt.Printf("List for Go routine# %d: ", myID)
	for _, v := range elements {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	mu.Unlock()
	sort.Ints(elements)
}

func mergeSort(a1, a2 []int) []int {
	result := []int{}
	for len(a1) > 0 || len(a2) > 0 {
		if len(a1) > 0 && len(a2) == 0 {
			result = append(result, a1[0])
			a1 = a1[1:]
			continue
		}
		if len(a1) == 0 && len(a2) > 0 {
			result = append(result, a2[0])
			a2 = a2[1:]
			continue
		}
		if a1[0] <= a2[0] {
			result = append(result, a1[0])
			a1 = a1[1:]
			continue
		}
		result = append(result, a2[0])
		a2 = a2[1:]
	}
	return result
}

func main() {
	var wg sync.WaitGroup
	kbd := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter at least 4 space separated INTs to sort >")
	input, _ := kbd.ReadString('\n')
	input = strings.TrimSpace(input)
	ints := strings.Split(input, " ")
	if len(ints) < 4 {
		fmt.Printf("Not enough space separated INTs\n")
		return
	}
	e := []int{}
	for i, v := range ints {
		j, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Bad INT %s at %d - %s\n", v, i, err)
			return
		}
		e = append(e, j)
	}

	var numPer int
	numPer = len(e) / 4
	leftover := len(e) % 4
	a1 := e[0 : numPer-1+1]
	a2 := e[numPer : numPer*2-1+1]
	a3 := e[numPer*2 : numPer*3-1+1]
	a4 := e[numPer*3 : numPer*4-1+1+leftover]

	wg.Add(1)
	go sorter(0, a1, &wg)
	wg.Add(1)
	go sorter(1, a2, &wg)
	wg.Add(1)
	go sorter(2, a3, &wg)
	wg.Add(1)
	go sorter(3, a4, &wg)
	wg.Wait()

	result := mergeSort(a1, a2)
	result = mergeSort(result, a3)
	result = mergeSort(result, a4)
	dumper("Final result", result)
}

func dumper(name string, a []int) {
	fmt.Printf("%s: ", name)
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
