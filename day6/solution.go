package day6

import (
	"fmt"
	"strconv"
	"strings"
)

const days = 80
const resetValue = 6
const newValue = 8

type Lanternfish = int

func Solution1(){
	lanternFish := initFish()

	for i := 0; i < days; i++ {
		var newFish []Lanternfish
		for i, fish := range lanternFish {
			if fish == 0 {
				lanternFish[i] = resetValue
				newFish = append(newFish, newValue)
			} else {
				lanternFish[i] -= 1
			}
		}

		lanternFish = append(lanternFish, newFish...)
	}

	fmt.Println("How many lanternfish would there be after 80 days?")
	fmt.Printf("%d\n", len(lanternFish))
}

func initFish() []Lanternfish {
	values := strings.Split(input, ",")

	lanternFish := make([]Lanternfish, len(values))
	for i, value := range values {
		fish, err := strconv.Atoi(value)
		if err != nil {
			panic("could not convert input to number " + value)
		}
		lanternFish[i] = fish
	}

	return lanternFish
}
