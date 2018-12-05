package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

import . "github.com/ahmetb/go-linq"

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	lines := strings.Split(input, "\n")

	frequenciesSeen := []int{0}
	frequency := 0
	repeatedFrequency := 0
	repeatedFrequencyFound := false
	for !repeatedFrequencyFound {
		for _, line := range lines {
			changeInFrequency, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			frequency += changeInFrequency
			seenBefore := From(frequenciesSeen).Contains(frequency)
			if seenBefore {
				repeatedFrequency = frequency
				repeatedFrequencyFound = true
				break
			}

			frequenciesSeen = append(frequenciesSeen, frequency)
		}
	}

	if repeatedFrequencyFound {
		log.Print(repeatedFrequency)
	} else {
		log.Print("Not found")
	}

}
