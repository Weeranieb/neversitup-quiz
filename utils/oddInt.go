package utils

func OddInt(numbers []int) int {
	numberMap := make(map[int]int)
	for _, n := range numbers {
		numberMap[n]++
	}

	for k, v := range numberMap {
		if v%2 != 0 {
			return k
		}
	}

	// if no odd number found, return 0
	return 0
}
