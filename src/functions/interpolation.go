package functions

func Interpolation(toSearch []int, toFind int) int {
	len := len(toSearch)
	low := 0
	high := len - 1

	for low <= high && toFind >= toSearch[low] && toFind <= toSearch[high] {
		if low == high {
			if toSearch[low] == toFind {
				return low
			}
			return -1
		}

		pos := low + int((high - low) / (toSearch[high] - toSearch[low])) * (toFind - toSearch[low])
		if toSearch[pos] == toFind {
			return pos
		}

		if toSearch[pos] < toFind {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}
