package main

import (
	"fmt"
	"input"
	"utils"
)

func main() {
	if !input.CheckExists(3) {
		input.Fetch(3)
	} else {
		fmt.Println("Input cached")
	}
	lines := utils.GetLines("./day3/input.txt")
	var result int32
	lines_length := len(lines)
	var lastIndex = 0
	for lineIndex, line := range lines {
		fmt.Println("indx", lineIndex)

		if lineIndex == lines_length-2 {
			fmt.Println("lines", lines_length, lineIndex)
			break
		}
		if lineIndex != lastIndex+1 && lineIndex != 0 {
			fmt.Println("skipping", lineIndex, lastIndex)
			continue
		}
		second_line := lines[lineIndex+1]
		third_line := lines[lineIndex+2]
		lastIndex = lineIndex + 2

		var tre_lines [3]map[rune]bool
		for num, line := range []string{line, second_line, third_line} {
			length := len(line)
			s := make(map[rune]bool, length)

			for _, c := range line {
				s[c] = true
			}
			tre_lines[num] = s
		}

		s_intersection := map[rune]bool{}
		s1, s2, s3 := tre_lines[0], tre_lines[1], tre_lines[2]
		for k, _ := range s1 {
			if s2[k] && s3[k] {
				s_intersection[k] = true
			}
		}

		for value := range s_intersection {
			if value != 0 {
				if value < 97 {
					// fmt.Println(string(value), value-38)
					result += value - 38
				} else {
					// fmt.Println(string(value), value-96)
					result += value - 96
				}
			}
			break
		}

	}
	fmt.Println("asnwer", result)

}
