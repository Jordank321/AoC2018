package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)

	lines := strings.Split(input, "\n")

	frequency := 0
	for _, line := range lines {
		changeInFrequency, err := strconv.Atoi(line)
		frequency += changeInFrequency

		if err != nil{
			panic(err)
		}
	}

	log.Print(frequency)
}
