package filter

import (
	"strings"

	"github.com/matheus-osorio/go-term-solver/src/entry"
)

type WordFilter struct {
	Rules []Rule
}

func (filter *WordFilter) CreateFilter(filterList []entry.Filter) {
	wordMap := map[rune]Rule{}

	for index, rule := range filterList {
		letter := rune(rule.Letter[0])
		current, exists := wordMap[letter]
		if !exists {
			current.Letter = letter
		}

		switch rule.Status {
		case entry.GREEN:
			current.In = append(current.In, index)
			current.NumberOfTimes++

		case entry.YELLOW:
			current.NotIn = append(current.NotIn, index)
			current.NumberOfTimes++

		case entry.GREY:
			current.Exact = true
		}

		wordMap[letter] = current
	}

	for _, value := range wordMap {
		filter.Rules = append(filter.Rules, value)
	}
}

func (filter WordFilter) IsWordValid(word string) bool {
	for _, rule := range filter.Rules {
		letter := rule.Letter
		for _, index := range rule.In {
			if word[index] != byte(letter) {
				return false
			}
		}

		for _, index := range rule.NotIn {
			if word[index] == byte(letter) {
				return false
			}

			strings.Contains(word, string(letter))
		}

		appearances := strings.Count(word, string(letter))

		if appearances < rule.NumberOfTimes || (appearances != rule.NumberOfTimes && rule.Exact) {
			return false
		}

	}

	return true
}

func (filter *WordFilter) FilterWords(wordList []string) (filteredWords []string) {
	for _, word := range wordList {
		if filter.IsWordValid(word) {
			filteredWords = append(filteredWords, word)
		}
	}

	return
}
