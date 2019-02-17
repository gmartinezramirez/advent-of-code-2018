package main

import "testing"

func TestShouldReturnTheCountOfAGivenLetter(t *testing.T) {
	var expected int = 2
	var result int = countLetterOfAString("Hello", "e")

	if result != expected {
		t.Errorf("got '%d' want '%d'", result, expected)
	}

}
