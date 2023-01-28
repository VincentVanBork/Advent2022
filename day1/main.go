package main

import (
	"fmt"
	"input"
)

func main() {
	if !input.CheckExists() {
		input.Fetch()
	} else {
		fmt.Println("Input cached")
	}

}
