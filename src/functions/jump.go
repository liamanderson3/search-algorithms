package functions

import (
	"math"
)

func Jump(to_search []int, to_find int) int {
	len := len(to_search)
	base_step := int(math.Sqrt(float64(len)))
	step := base_step
	prev := 0
	for to_search[Min(step, len)-1] < to_find {
		prev = step
		step += base_step
		if prev >= to_find {
			return -1
		}
	}
	
	for to_search[prev] < to_find {
		prev += 1
		if prev == Min(step, len) {
			return -1
		}
	}

	if to_search[prev] == to_find {
		return prev
	}

	return -1
}
