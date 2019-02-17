package main

import (
	"fmt"
	"strings"
)

func countLetterOfAString(word, letter string) int {
	return strings.Count(word, letter)
}

func main() {
	fmt.Println(countLetterOfAString("Hello", "e"))
}
