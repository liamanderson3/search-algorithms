package main

import (
	"fmt"
	"functions"
)

func main() {
	// to_sort := []int{15, 2, 1, 123, 272, 72, 727, 17, 8222, 1, 17, 25, 4}
	sortedHaystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	toFind := 72
	toPrint := functions.Interpolation(sortedHaystack, toFind)
	fmt.Println(toPrint)
}