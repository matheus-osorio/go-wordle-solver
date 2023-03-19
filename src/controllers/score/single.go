package score

type SingleScoreSystem struct {
	PositionPoints []map[rune]float64
	LetterPoints   map[rune]float64
	Points         ScoreList
	WordList       []string
	WordSize       int
	TotalWords     int
}

func (score SingleScoreSystem) getWordScore(word string) float64 {
	var greenScore, yellowScore, greyScore float64

	for index, letter := range word {

		totalWords := float64(score.TotalWords)
		currentPositionScore := score.PositionPoints[index][letter]
		currentWordScore := score.LetterPoints[letter]
		wordsNotInPosition := (currentWordScore - currentPositionScore)
		wordsWithoutLetter := (totalWords - currentPositionScore)

		greenProbability := currentPositionScore / totalWords
		yellowProbability := wordsNotInPosition / totalWords
		greyProbability := wordsWithoutLetter / totalWords

		greenScore += (totalWords - currentPositionScore) * greenProbability
		yellowScore += (totalWords - wordsNotInPosition) * yellowProbability
		greyScore += (totalWords - wordsWithoutLetter) * greyProbability
	}
	return greenScore + yellowScore + greyScore
}

func (score *SingleScoreSystem) createScoreTable() {
	score.PositionPoints = make([]map[rune]float64, 0)
	score.LetterPoints = make(map[rune]float64, 0)

	for i := 0; i < score.WordSize; i++ {
		newMap := make(map[rune]float64)
		score.PositionPoints = append(score.PositionPoints, newMap)
	}
	for _, word := range score.WordList {
		for index, letter := range word {
			score.PositionPoints[index][letter]++
		}
		letterSet := createSet(word)

		for _, letter := range letterSet {
			score.LetterPoints[letter]++
		}
	}
}

func (score *SingleScoreSystem) setScores() {
	score.Points = ScoreList{}
	for _, word := range score.WordList {
		score.Points = append(score.Points, Score{
			Word:   word,
			Points: score.getWordScore(word),
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
