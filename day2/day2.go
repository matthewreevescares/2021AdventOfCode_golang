package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//import "strconv"

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
	}

	//create a int slice with that has a length and capacity of that data string slice
	fmt.Println(strings.Split(data[1], " ")[0])
	ints := []int{}
	direction := []string{}

	for i := 0; i < len(data); i++ {
		s := strings.Split(data[i], " ")
		direction = append(direction, s[0])
		a, _ := strconv.Atoi(s[1])
		ints = append(ints, a)
	}

	//Part 1
	depth, horizonal := 0, 0
	if len(direction) == len(ints) {
		for i := range direction {
			if direction[i] == "forward" {
				horizonal += ints[i]
			}
			if direction[i] == "down" {
				depth += ints[i]
			}
			if direction[i] == "up" {
				depth -= ints[i]
			}
		}
	}
	fmt.Println(depth)
	fmt.Println(horizonal)
	fmt.Println(depth * horizonal)

	//Part 2 with aim
	depth, horizonal, aim := 0, 0, 0
	if len(direction) == len(ints) {
		for i := range direction {
			if direction[i] == "forward" {
				horizonal += ints[i]
				depth += (aim * ints[i])
			}
			if direction[i] == "down" {
				aim += ints[i]
			}
			if direction[i] == "up" {
				aim -= ints[i]
			}
		}
	}

	fmt.Println(depth)
	fmt.Println(horizonal)
	fmt.Println(depth * horizonal)

}
