package functions

func Binary(to_search []int, left int, right int, to_find int) int {
	if right >= 1 {
		mid := left + (right - left)/2
		if to_search[mid] == to_find {
			return mid
		} else if to_search[mid] > to_find {
			return Binary(to_search, left, mid - 1, to_find)
		} else {
			return Binary(to_search, mid + 1, right, to_find)
		}
	}
	return -1
}
