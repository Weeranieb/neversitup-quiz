package utils

import "fmt"

func Permutation(letter string) []string {
	// count all characters
	countMap := make(map[rune]int)
	for _, c := range letter {
		countMap[c]++
	}

	// create all possible permutations
	var result []string
	permutationHelper(countMap, "", len(letter), &result)
	return result
}

func permutationHelper(countMap map[rune]int, prefix string, remaining int, result *[]string) {
	if remaining == 0 {
		*result = append(*result, prefix)
		return
	}

	for c, count := range countMap {
		if count > 0 {
			countMap[c]--
			fmt.Println(prefix + string(c))
			permutationHelper(countMap, prefix+string(c), remaining-1, result)
			countMap[c]++
		}
	}
}
