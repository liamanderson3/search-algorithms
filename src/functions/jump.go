package functions

import (
	"math"
)

func Jump(toSearch []int, toFind int) int {
	len := len(toSearch)
	baseStep := int(math.Sqrt(float64(len)))
	step := baseStep
	prev := 0
	for toSearch[Min(step, len)-1] < toFind {
		prev = step
		step += baseStep
		if prev >= toFind {
			return -1
		}
	}

	for toSearch[prev] < toFind {
		prev++
		if prev == Min(step, len) {
			return -1
		}
	}

	if toSearch[prev] == toFind {
		return prev
	}

	return -1
}
