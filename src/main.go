package main

import (
	"fmt"
	"functions"
)

func main() {
	// to_sort := []int{15, 2, 1, 123, 272, 72, 727, 17, 8222, 1, 17, 25, 4}
	sorted_haystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	to_find := 72
	to_print := functions.Interpolation(sorted_haystack, to_find)
	fmt.Println(to_print)
}