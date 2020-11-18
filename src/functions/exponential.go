package functions

func Exponential(toSearch []int, toFind int) int {
	if toSearch[0] == toFind {
		return 0
	}
	len := len(toSearch)
	i := 1
	for i < len && toSearch[i] <= toFind {
		i *= 2
	}
	return Binary(toSearch, i / 2, Min(i, len-1), toFind)
}
