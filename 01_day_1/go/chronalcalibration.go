package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkForErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	checkForErrors(err)
	defer file.Close()

	currentFrequency := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		changeFrequency, _ := strconv.Atoi(scanner.Text())
		currentFrequency += changeFrequency
	}
	fmt.Println(currentFrequency)
}
