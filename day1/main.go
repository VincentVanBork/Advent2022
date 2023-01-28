package main

import (
	"fmt"
	"input"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup
	elvesChannel := make(chan int, len(elves))
	for _, elf := range elves {
		wg.Add(1)
		go utils.SumElements(elf, elvesChannel, &wg)
	}

	go func(channel chan int, wg *sync.WaitGroup) {
		wg.Wait()
		close(channel)
	}(elvesChannel, &wg)

	maxCalories := 0
	for elfCalorie := range elvesChannel {
		if elfCalorie > maxCalories {
			maxCalories = elfCalorie
		}
	}
	fmt.Println("Final answer:", maxCalories)
}
