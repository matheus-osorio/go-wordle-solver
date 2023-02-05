package score

import (
	"math"
	"strings"
	"sync"
)

type SingleScoreSystem struct {
	ScorePoints []map[rune]float64
	Points      ScoreList
	WordList    []string
	WordSize    int
}

func (score SingleScoreSystem) getWordScore(word string) (letterScore, positionScore float64) {
	passedWords := map[rune]bool{}
	for index, letter := range word {
		positionScore += score.ScorePoints[index][letter]
		if !passedWords[letter] {
			specificWordScore := 0.0
			for i := 0; i < score.WordSize; i++ {

				if word[i] != byte(letter) {
					specificWordScore += score.ScorePoints[i][letter]
				}
			}
			passedWords[letter] = true

			if specificWordScore > 0 {
				specificWordScore /= math.Pow(float64(score.WordSize-strings.Count(word, string(letter))), 2)
				letterScore += specificWordScore
			}
		}

	}
	return
}

func (score *SingleScoreSystem) createScoreTable() {
	score.ScorePoints = make([]map[rune]float64, 0)
	for i := 0; i < score.WordSize; i++ {
		newMap := make(map[rune]float64)
		score.ScorePoints = append(score.ScorePoints, newMap)
	}
	for _, word := range score.WordList {
		for index, letter := range word {
			score.ScorePoints[index][letter]++
		}
	}
}

func (score *SingleScoreSystem) setScores() {
	score.Points = ScoreList{}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(score.WordList))
	for _, word := range score.WordList {
		letter, position := score.getWordScore(word)
		score.Points = append(score.Points, Score{
			Word:     word,
			Letter:   letter,
			Position: position,
		})
	}
}

// Scores all the words and gives the score back
func (score *SingleScoreSystem) GetBestWords() []ScoreList {
	score.createScoreTable()
	score.setScores()
	score.Points = score.Points.SortByScore()
	scoreList := []ScoreList{
		score.Points,
	}

	return scoreList
}
