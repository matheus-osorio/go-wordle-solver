package score

func createSet(word string) []rune {
	mapSet := map[rune]bool{}

	for _, letter := range word {
		mapSet[letter] = true
	}

	result := []rune{}

	for key := range mapSet {
		result = append(result, key)
	}

	return result
}
