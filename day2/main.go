package main

import (
	"fmt"
	"input"
	"strings"
	"utils"
)

var options = map[string]int{"A": 0, "B": 1, "C": 2, "X": 0, "Y": 1, "Z": 2}

func getOptions(opp string, player string) (int, int) {
	// X loose y draw z win
	// spaghetti code here we go
	option := 0
	switch player {
	case "X":
		switch opp {
		case "A":
			option = 2
			break
		case "B":
			option = 0
			break
		case "C":
			option = 1
			break
		}
		break
	case "Y":
		switch opp {
		case "A":
			option = 0
			break
		case "B":
			option = 1
			break
		case "C":
			option = 2
			break
		}
		break
	case "Z":
		switch opp {
		case "A":
			option = 1
			break
		case "B":
			option = 2
			break
		case "C":
			option = 0
			break
		}
		break
	}
	// row and column
	return option, options[opp]
}

func getRules() [][]int {
	playValues := []int{1, 2, 3}
	rulesMatrix := [][]int{
		{3, 0, 6},
		{6, 3, 0},
		{0, 6, 3},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < len(rulesMatrix[i]); j++ {
			rulesMatrix[i][j] += playValues[i]
		}
	}
	return rulesMatrix
}

func play(opponent string, player string) int {
	rulesMatrix := getRules()
	// fmt.Println(rulesMatrix)
	i, j := getOptions(opponent, player)
	return rulesMatrix[i][j]
}
func main() {
	if !input.CheckExists(2) {
		input.Fetch(2)
	} else {
		fmt.Println("Input cached")
	}
	lines := utils.GetLines("./day2/input.txt")
	fmt.Println(lines[0])
	var sum int
	for _, line := range lines {
		values := strings.Split(line, " ")
		opponent, you := values[0], values[1]
		outcome := play(opponent, you)
		sum += outcome
	}
	fmt.Println("result", sum)
}
