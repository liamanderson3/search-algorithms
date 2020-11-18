package functions

func Linear(param []int, toFind int) int {
	for i := 0; i < len(param); i++ {
		if param[i] == toFind {
			return i
		}
	}
	return -1
}
