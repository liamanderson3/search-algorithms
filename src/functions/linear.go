package functions

func Linear(param []int, to_find int) int {
	for i := 0; i < len(param); i++ {
		if param[i] == to_find {
			return i
		}
	}
	return -1
}
