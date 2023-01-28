package main

import (
	"fmt"
	"input"
)

func main() {
	if !input.CheckExists(2) {
		input.Fetch(2)
	} else {
		fmt.Println("Input cached")
	}
}
