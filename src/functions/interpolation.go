package functions

func Interpolation(to_search []int, to_find int) int {
	len := len(to_search)
	low := 0
	high := len - 1

	for low <= high && to_find >= to_search[low] && to_find <= to_search[high] {
		if low == high {
			if to_search[low] == to_find {
				return low
			}
			return -1
		}

		pos := low + int((high - low) / (to_search[high] - to_search[low])) * (to_find - to_search[low])
		if to_search[pos] == to_find {
			return pos
		}

		if to_search[pos] < to_find {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}
