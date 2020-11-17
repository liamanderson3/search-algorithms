package main

import (
	"fmt"
	"functions"
)

func main() {
	to_sort := []int{15, 2, 1, 123, 272, 72, 727, 17, 8222, 1, 17, 25, 4}
	to_find := 72
	to_print := functions.Linear(to_sort, to_find)
	fmt.Println(to_print)
}