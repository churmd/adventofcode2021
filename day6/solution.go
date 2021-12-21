package day6

import (
	"fmt"
	"strconv"
	"strings"
)

const resetValue = 6
const newValue = 8

type Lanternfish = int

type LanternfishManager = map[Lanternfish]int

func Solution1() {
	lanternFish := initFish()

	for i := 0; i < 80; i++ {
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

func Solution2() {
	manager := NewLaternfishManager()

	for i := 0; i < 256; i++ {
		updatedManager := LanternfishManager{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

		updatedManager[8] = manager[0]
		updatedManager[7] = manager[8]
		updatedManager[6] = manager[7] + manager[0]
		updatedManager[5] = manager[6]
		updatedManager[4] = manager[5]
		updatedManager[3] = manager[4]
		updatedManager[2] = manager[3]
		updatedManager[1] = manager[2]
		updatedManager[0] = manager[1]

		manager = updatedManager
	}

	fmt.Println("How many lanternfish would there be after 256 days?")
	numFish := manager[0] + manager[1] + manager[2] + manager[3] + manager[4] + manager[5] + manager[6] + manager[7] + manager[8]
	fmt.Printf("%d\n", numFish)
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

func NewLaternfishManager() LanternfishManager {
	manager := LanternfishManager{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	values := strings.Split(input, ",")
	for _, value := range values {
		fish, err := strconv.Atoi(value)
		if err != nil {
			panic("could not convert input to number " + value)
		}
		manager[fish] += 1
	}

	return manager
}
