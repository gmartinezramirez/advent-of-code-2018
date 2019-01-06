package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func checkForErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func readFrequencyFromFile(fileHandler io.Reader) int {
	currentFrequency := 0

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		changeFrequency, _ := strconv.Atoi(scanner.Text())
		currentFrequency += changeFrequency
	}

	return currentFrequency
}

func main() {
	fileHandler, err := os.Open("input.txt")
	checkForErrors(err)
	defer fileHandler.Close()
	currentFrequency := readFrequencyFromFile(fileHandler)
	fmt.Println(currentFrequency)
}
