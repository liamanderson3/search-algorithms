package functions

import (
	"unsafe"
)

func Linear(param []int, to_find int) int {
	index := -1
	for i := 0; i < int(unsafe.Sizeof(param)); i++ {
		if param[i] == to_find {
			return i
		}
	}
	return index
}
