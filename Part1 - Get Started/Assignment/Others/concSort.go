/*
 * concSort.go
 * Author: LTrevino
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 *
 * Please note that I like using semicolons -- valid in Go.
 */
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const N_THREADS int = 4;

/*
 Author: LTrevino. Copyright (C) 2107 Louis Trevino.
 * Using this function, since:
 *   - A single input could contain spaces (e.g. "Jane Doe")
 *   - fmt.Scan() is ignoring characters after 1st space, e.g. in "I d skd a efju N" only "I" will be read.
 *   - fmt.Scanln() throwing warning with "Enter" (newline) or at least if running Windows 10 Home Edition.
*/
func acceptLn() (usrInput string, err error) {
	reader := bufio.NewReader(os.Stdin)
	if usrInput, err = reader.ReadString('\n'); err != nil {
		return usrInput, err
	} else {
		//usrInput = strings.Replace(usrInput, "\r\n", "", -1)
		re := regexp.MustCompile("\r?\n")
		usrInput = re.ReplaceAllString(usrInput, "")
	}
	//fmt.Printf("usrInput: %s \n", usrInput)
	return usrInput, err
}

func parseIntegers(intStrList string) (intSlice []int) {
	intStrings := strings.Split(intStrList, " ");
	intSlice = make([]int, 0, 20);
	for _, str := range intStrings {
		nn, err := strconv.Atoi(str);
		if (err!=nil) {
			fmt.Printf("Number '%s' could not be parsed -- ignored. \n")
		} else {
			intSlice = append(intSlice, nn);
		}
	}
	return intSlice;
}

func sortChunk(ch chan string, chunk []int, threadId int, fromIdx int, toIdx int) {
	fmt.Printf("Thread %d <-- Raw chunk    [%d:%d] : %v \n", threadId, fromIdx, toIdx, chunk);
	sort.Ints(chunk);
	fmt.Printf("Thread %d --> Sorted chunk [%d:%d] : %v \n", threadId, fromIdx, toIdx, chunk);
	ch <- "done";
}

func splitSort(ch chan string, numList []int) {
	var chunkSize int = int(math.Ceil( float64(len(numList)) / float64(N_THREADS) ));
	// fmt.Printf("Chunk size: %d \n", chunkSize);
	var thCount = 0;
	for iPage := 0; iPage < N_THREADS; iPage++ {
		iFrom := iPage * chunkSize;
		if iFrom >= len(numList) {
			break;
		}
		thCount++;
		iTo := iFrom + chunkSize;
		if (iTo > len(numList)) {
			iTo = len(numList)
		}
		iChunk := numList[iFrom:iTo];
		go sortChunk(ch, iChunk, iPage, iFrom, iTo);
	}
	for idx:=0; idx < thCount; idx++ {
		<- ch;
	}
	// fmt.Printf("Merged list:      %d \n", numList);
	sort.Ints(numList);
	fmt.Printf("Sorted full list: %d \n", numList);
}

func main() {
	ch := make(chan string);
	fmt.Println("Welcome to concSort.go");
	//var numList = []int {20,19,18,17, 16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1};
	fmt.Println("Enter integer numbers: separated by space (e.g. \"7 6 5 4 3 2 1\")");
	usrInput, err := acceptLn();
	if (err!=nil) {fmt.Printf("Error reading numbers. %s\n", err)}
	//fmt.Println("usr input: ", usrInput)
	numList := parseIntegers(usrInput);
	fmt.Printf("Original list:     %v \n" , numList);
	splitSort(ch, numList);
}

