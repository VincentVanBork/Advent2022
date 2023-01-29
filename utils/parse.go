package utils

import (
	"bufio"
	"os"
)

func GetLines(filePath string) []string {
	file, err := os.Open(filePath)
	Check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	defer file.Close()

	return fileLines
}
