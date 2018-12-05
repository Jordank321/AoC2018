package main

import (
	"io/ioutil"
	"strings"

	. "github.com/ahmetb/go-linq"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	lines := strings.Split(input, "\n")

	boxesWith2SameLetters := countIDsWithXLetters(2, lines)
	boxesWith3SameLetters := countIDsWithXLetters(3, lines)

	print(boxesWith2SameLetters * boxesWith3SameLetters)

}

func countIDsWithXLetters(letterFrequency int, ids []string) int {
	return From(ids).CountWith(func(id interface{}) bool {
		hasAnyXOccurrencesOfALetter := From(strings.Split(id.(string), "")).Distinct().AnyWith(func(searchChar interface{}) bool {
			occurrencesOfSearchChar := From(strings.Split(id.(string), "")).CountWith(func(char interface{}) bool {
				return char.(string) == searchChar.(string)
			})
			return occurrencesOfSearchChar == letterFrequency
		})
		return hasAnyXOccurrencesOfALetter
	})
}
