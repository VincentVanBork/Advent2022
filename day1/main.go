package main

import (
	"fmt"
	"input"
	"sort"
	"strconv"
	"sync"
	"utils"
)

func main() {
	if !input.CheckExists(1) {
		input.Fetch(1)
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
	var sumElves []int
	for elfCalorie := range elvesChannel {
		sumElves = append(sumElves, elfCalorie)
	}
	sort.Slice(sumElves, func(i, j int) bool {
		return sumElves[i] > sumElves[j]
	})

	for _, v := range sumElves[:3] {
		maxCalories += v
	}
	fmt.Println("Final answer:", maxCalories)
}
