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

func (solver WordleSolver) GetFullWordList() []score.ScoreList {
	words := solver.Getter.GetWords()
	solver.Score = score.CreateFactory([][]string{words}, solver.WordSize)

	return solver.Score.GetBestWords()
}

func (solver WordleSolver) FilterWords(entry entry.FilterListBodyObject) []score.ScoreList {
	newWordList := [][]string{}

	for _, ruleList := range entry.Rules {
		solver.Filter.CreateFilter(ruleList.Filter)
		newWordList = append(newWordList, solver.Filter.FilterWords(ruleList.Words))
	}

	solver.Score = score.CreateFactory(newWordList, solver.WordSize)

	return solver.Score.GetBestWords()
}
