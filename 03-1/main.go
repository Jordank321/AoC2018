package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	. "github.com/ahmetb/go-linq"
)

type claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	expression := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")

	claims := []claim{}
	From(lines).Select(func(line interface{}) interface{} {
		matches := expression.FindStringSubmatch(line.(string))
		claim := claim{}
		claim.ID, err = strconv.Atoi(matches[1])
		if err != nil {
			log.Panic(err)
		}
		claim.X, err = strconv.Atoi(matches[2])
		if err != nil {
			log.Panic(err)
		}
		claim.Y, err = strconv.Atoi(matches[3])
		if err != nil {
			log.Panic(err)
		}
		claim.Width, err = strconv.Atoi(matches[4])
		if err != nil {
			log.Panic(err)
		}
		claim.Height, err = strconv.Atoi(matches[5])
		if err != nil {
			log.Panic(err)
		}
		return claim
	}).ToSlice(&claims)

	claimsOnFabric := make(map[string]int)

	for _, claim := range claims {
		for x := claim.X; x < claim.Width+claim.X; x++ {
			for y := claim.Y; y < claim.Height+claim.Y; y++ {
				key := strconv.Itoa(x) + "," + strconv.Itoa(y)
				claimsOnFabric[key]++
			}
		}
	}

	overlaps := 0
	for _, value := range claimsOnFabric {
		if value > 1 {
			overlaps++
		}
	}

	log.Print(overlaps)
}
