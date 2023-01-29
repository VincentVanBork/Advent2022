package main

import (
	"fmt"
	"input"
)

func main() {
	if !input.CheckExists(3) {
		input.Fetch(3)
	} else {
		fmt.Println("Input cached")
	}
}
