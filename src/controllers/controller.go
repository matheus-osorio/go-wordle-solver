package controllers

import (
	"github.com/matheus-osorio/go-term-solver/src/controllers/filter"
	"github.com/matheus-osorio/go-term-solver/src/controllers/getters"
	"github.com/matheus-osorio/go-term-solver/src/controllers/score"
	"github.com/matheus-osorio/go-term-solver/src/entry"
)

type WordleSolver struct {
	Getter   getters.GetterInterface
	Score    score.ScoreInterface
	Filter   filter.WordFilter
	WordSize int
}

// Gets an unfiltered word list
func (solver WordleSolver) GetFullWordList() []score.ScoreList {
	words := solver.Getter.GetWords()
	solver.Score = score.CreateScore([][]string{words}, solver.WordSize)

	return solver.Score.GetBestWords()
}

// Receives an entry object of lists to filter
func (solver WordleSolver) FilterWords(entry entry.FilterListBodyObject) []score.ScoreList {
	newWordList := [][]string{}

	for _, ruleList := range entry.Rules {
		solver.Filter.CreateFilter(ruleList.Filter)
		newWordList = append(newWordList, solver.Filter.FilterWords(ruleList.Words))
	}

	solver.Score = score.CreateScore(newWordList, solver.WordSize)

	return solver.Score.GetBestWords()
}
