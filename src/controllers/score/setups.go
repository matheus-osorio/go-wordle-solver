package score

type ScoreInterface interface {
	GetBestWords() []ScoreList
}

type Score struct {
	Word   string
	Points float64
}

type ScoreList []Score

// Sorts without knowing method
func (score ScoreList) SortByScore() ScoreList {
	return score.quickSort(score)
}

// Uses quick sort to get the best words
func (score ScoreList) quickSort(points ScoreList) ScoreList {
	if len(points) < 2 {
		return points
	}
	middle := int(len(points) / 2)
	middleScore := points[middle]

	rightSide := ScoreList{}
	leftSide := ScoreList{}

	for index, value := range points {
		comparedScores := score.compareScores(middleScore, value)

		if comparedScores == 0 {
			if index != middle {
				rightSide = append(rightSide, value)
			}
			continue
		}

		if comparedScores > 0 {
			rightSide = append(rightSide, value)
			continue
		}

		leftSide = append(leftSide, value)
	}

	leftSide = score.quickSort(leftSide)
	rightSide = score.quickSort(rightSide)

	orderedArr := append(leftSide, middleScore)
	orderedArr = append(orderedArr, rightSide...)
	return orderedArr
}

// Compares the scores
func (score ScoreList) compareScores(scoreObj1, scoreObj2 Score) int {
	score1 := scoreObj1.Points
	score2 := scoreObj2.Points

	if score1 > score2 {
		return 1
	}

	if score1 < score2 {
		return -1
	}

	return 0
}

// Gets only the top words (currently unnutilized)
func (score ScoreList) GetTop(number int) ScoreList {
	if len(score) <= number {
		return score
	}

	return score[:number]
}
