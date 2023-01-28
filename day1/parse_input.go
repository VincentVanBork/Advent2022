package main

import (
	"bufio"
	"os"
	"utils"
)

func GetLines(filePath string) []string {
	file, err := os.Open(filePath)
	utils.Check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	// for _, line := range fileLines {
	// 	fmt.Println(line)
	// }
	return fileLines
}
