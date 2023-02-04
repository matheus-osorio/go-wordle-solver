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

func (score SingleScoreSystem) GetWordScore(word string) (letterScore, positionScore float64) {
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
			specificWordScore /= math.Pow(float64(score.WordSize-strings.Count(word, string(letter))), 2)
			letterScore += specificWordScore
		}

	}
	return
}

func (score *SingleScoreSystem) CreateScoreTable() {
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

func (score *SingleScoreSystem) SetScores() {
	score.Points = ScoreList{}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(score.WordList))
	for _, word := range score.WordList {
		letter, position := score.GetWordScore(word)
		score.Points = append(score.Points, Score{
			Word:     word,
			Letter:   letter,
			Position: position,
		})
	}
}

func (score *SingleScoreSystem) GetBestWords() []ScoreList {
	score.CreateScoreTable()
	score.SetScores()
	score.Points = score.Points.SortByScore()
	scoreList := []ScoreList{
		score.Points,
	}

	return scoreList
}
