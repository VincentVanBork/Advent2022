package main

import (
	"fmt"
	"input"
	"strconv"
	"utils"
)

func main() {
	if !input.CheckExists() {
		input.Fetch()
	} else {
		fmt.Println("Input cached")
	}
	var elves [][]int
	lines := GetLines("./day1/input.txt")
	var elf []int
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf)
			elf = nil
		} else {
			val, err := strconv.Atoi(line)
			utils.Check(err)
			elf = append(elf, val)
		}
	}
	fmt.Println(elves[0])

}
