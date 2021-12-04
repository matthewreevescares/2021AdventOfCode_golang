package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Open file from the OS
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	//scan the opened file
	scanner := bufio.NewScanner(file)
	//create a string slice for the line data to go into
	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		//Foreach line of data(string) append to the data slice
		data = append(data, line)
		fmt.Println(line)
	}

	fmt.Println(data)
	//create a int slice with that has a length of that data string slice
	ints := make([]int, len(data))

	//Foreach value in in string data slice convert and append to the int slice.
	for i, s := range data {
		ints[i], _ = strconv.Atoi(s)
	}

	//Create a total int slice to keep results that match the requirment of the algo
	total := []int{}
	for i := 0; i < len(ints); i++ {
		ii := i - 1
		if ints[i] != ints[0] && ints[i] > (ints[ii]-1) {
			total = append(total, ints[i])
		}

	}
	fmt.Println("List of data that are larger than the previous measurement", total)
	fmt.Println("How many data are larger than the previous measurement?", len(total))
}
