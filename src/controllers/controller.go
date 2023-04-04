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

func (solver *WordleSolver) createScore(words [][]string, removedWords [][]string) {
	solver.score = score.CreateScore(words, removedWords, solver.wordSize)
}

func (solver WordleSolver) getBestWords() ([]score.ScoreList, []score.ScoreList) {
	return solver.score.GetBestWords()
}

func (solver WordleSolver) filterWords(rules []entry.ObjectRuleList) ([][]string, [][]string) {
	wordList := [][]string{}
	removedList := [][]string{}

	for _, ruleList := range rules {
		solver.filter.CreateFilter(ruleList.Filter)
		filteredWords, removedWords := solver.filter.FilterWords(ruleList.Words)
		wordList = append(wordList, filteredWords)
		removedWords = append(removedWords, ruleList.RemovedWords...)
		removedList = append(removedList, removedWords)
	}

	return wordList, removedList
}

// Gets an unfiltered word list
func (solver WordleSolver) GetFullWordList() ([]score.ScoreList, []score.ScoreList) {
	words := solver.getWords()

	solver.createScore([][]string{words}, [][]string{{}})

	return solver.getBestWords()
}

// Receives an entry object of lists to filter
func (solver WordleSolver) FilterWords(entry entry.FilterListBodyObject) ([]score.ScoreList, []score.ScoreList) {
	wordList, removedList := solver.filterWords(entry.Rules)

	solver.createScore(wordList, removedList)

	return solver.getBestWords()
}
