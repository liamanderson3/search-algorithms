package functions

import (
	"unsafe"
)

func Fibonacci(param []int, to_find int) int {
	a := unsafe.Sizeof(param)
	return int(a)
}
