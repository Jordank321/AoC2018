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

func (theClaim claim) String() string {
	return strconv.Itoa(theClaim.ID)
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
		fatalCheck(err)
		claim.Y, err = strconv.Atoi(matches[3])
		fatalCheck(err)
		claim.Width, err = strconv.Atoi(matches[4])
		fatalCheck(err)
		claim.Height, err = strconv.Atoi(matches[5])
		fatalCheck(err)
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

	for key, value := range claimsOnFabric {
		if value > 1 {
			strings := strings.Split(key, ",")
			x, err := strconv.Atoi(strings[0])
			fatalCheck(err)
			y, err := strconv.Atoi(strings[1])
			fatalCheck(err)
			From(claims).Where(func(claimValue interface{}) bool {
				theClaim := claimValue.(claim)
				return !(theClaim.X <= x && theClaim.X+theClaim.Width > x && theClaim.Y <= y && theClaim.Y+theClaim.Height > y)
			}).ToSlice(&claims)
		}
	}

	for _, theClaim := range claims {
		log.Printf("Remaining claim: %s", theClaim)
	}
}

func fatalCheck(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
