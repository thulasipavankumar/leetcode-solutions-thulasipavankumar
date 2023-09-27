package main

import "fmt"

func countBits(n int) []int {
	var bytes, toReturn []int
	maxDiv := 2

	for n >= 1 {
		//bytes = append(bytes)

		bytes = append(bytes, n%maxDiv)
		n = n / maxDiv
		//fmt.Print(" ",n/maxDiv,bytes)

	}
	bytes = append(bytes, 0)
	for i := len(bytes) - 1; i >= 0; i-- {
		toReturn = append(toReturn, bytes[i])
	}

	return toReturn
}

func main() {
	fmt.Println(countBits(10))
}
