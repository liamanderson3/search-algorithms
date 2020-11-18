package functions

func Exponential(to_search []int, to_find int) int {
	if to_search[0] == to_find {
		return 0
	}
	len := len(to_search)
	i := 1
	for i < len && to_search[i] <= to_find {
		i *= 2
	}
	return Binary(to_search, i / 2, Min(i, len-1), to_find)
}
