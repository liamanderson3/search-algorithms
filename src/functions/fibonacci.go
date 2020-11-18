package functions

func Fibonacci(toSearch []int, toFind int) int {
	len := len(toSearch)
	fibNm2 := 0
	fibNm1 := 1
	fibN := fibNm2 + fibNm1

	for fibN < len {
		fibNm2 = fibNm1
		fibNm1 = fibN
		fibN = fibNm2 + fibNm1
	}

	index, offset := 0, -1
	for fibN > 1 {
		index = Min(offset+fibNm2, toFind-1)
		if toSearch[index] < toFind {
			fibN = fibNm1
			fibNm1 = fibNm2
			fibNm2 = fibN - fibNm1
			offset = index
		} else if toSearch[index] > toFind {
			fibN = fibNm2
			fibNm1 = fibNm1 - fibNm2
			fibNm2 = fibN - fibNm1
		} else {
			return index
		}
	}

	if fibNm1 > 0 && toSearch[offset+1] == toFind {
		return offset + 1
	}

	return -1
}
