package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	lines := strings.Split(input, "\n")

	result := ""
	for i, originalLine := range lines {
		for j, compareLine := range lines {

			//Find common runes in before-current and current-after
			result = getCommonLetters(originalLine, compareLine)
			if len(result) == len(originalLine)-1 {
				log.Printf("Result %s from IDs %s and %s at index %d and %d", result, originalLine, compareLine, i, j)
				return
			}
		}
	}
}

func getCommonLetters(first string, second string) string {
	commonChars := []string{}
	for i, char := range first {
		charSecond := second[i]
		if char == rune(charSecond) {
			commonChars = append(commonChars, string(charSecond))
		}
	}
	return strings.Join(commonChars, "")
}
