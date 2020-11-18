package functions

func Binary(toSearch []int, left int, right int, toFind int) int {
	if right >= 1 {
		mid := left + (right - left)/2
		if toSearch[mid] == toFind {
			return mid
		} else if toSearch[mid] > toFind {
			return Binary(toSearch, left, mid - 1, toFind)
		} else {
			return Binary(toSearch, mid + 1, right, toFind)
		}
	}
	return -1
}
