package controllers

import (
	"github.com/matheus-osorio/go-term-solver/src/controllers/filter"
	"github.com/matheus-osorio/go-term-solver/src/controllers/getters"
	"github.com/matheus-osorio/go-term-solver/src/controllers/score"
	"github.com/matheus-osorio/go-term-solver/src/entry"
)

type WordleSolver struct {
	getter   getters.GetterInterface
	score    score.ScoreInterface
	filter   filter.WordFilter
	wordSize int
}

func (solver WordleSolver) getWords() []string {
	return solver.getter.GetWords()
}

func (solver *WordleSolver) createScore(words [][]string) {
	solver.score = score.CreateScore(words, solver.wordSize)
}

func (solver WordleSolver) getBestWords() []score.ScoreList {
	return solver.score.GetBestWords()
}

func (solver WordleSolver) filterWords(rules []entry.ObjectRuleList) [][]string {
	wordList := [][]string{}

	for _, ruleList := range rules {
		solver.filter.CreateFilter(ruleList.Filter)
		wordList = append(wordList, solver.filter.FilterWords(ruleList.Words))
	}

	return wordList
}

// Gets an unfiltered word list
func (solver WordleSolver) GetFullWordList() []score.ScoreList {
	words := solver.getWords()

	solver.createScore([][]string{words})

	return solver.getBestWords()
}

// Receives an entry object of lists to filter
func (solver WordleSolver) FilterWords(entry entry.FilterListBodyObject) []score.ScoreList {
	wordList := solver.filterWords(entry.Rules)

	solver.createScore(wordList)

	return solver.getBestWords()
}
