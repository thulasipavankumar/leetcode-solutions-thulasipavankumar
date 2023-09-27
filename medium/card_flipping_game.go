package main

import (
	"fmt"
	"math"
)

func flipgame(fronts []int, backs []int) int {
	// Cards which have same number on front and back
	mirrorCards := make(map[int]bool)

	for index, val := range fronts {
		if fronts[index] == backs[index] {
			mirrorCards[val] = true
		}
	}
	fmt.Println("Mirroed cards", mirrorCards)
	res := math.MaxInt32
	for index, val := range fronts {
		if val < res && !mirrorCards[val] {
			res = val
		}
		if backs[index] < res && !mirrorCards[backs[index]] {
			res = backs[index]
		}
	}

	if res == math.MaxInt32 {
		return 0
	} else {
		return res
	}
}

func main() {
	//fronts,backs := []int{1,2,4,4,7},[]int{1,3,4,1,3}

	fronts, backs := []int{1, 1}, []int{2, 2}

	fmt.Println(flipgame(fronts, backs))

}
