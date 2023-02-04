package score

type ScoreInterface interface {
	GetBestWords() []ScoreList
}

type Score struct {
	Word     string
	Letter   float64
	Position float64
}

type ScoreList []Score

func (score ScoreList) SortByScore() ScoreList {
	return score.QuickSort(score)
}

func (score ScoreList) QuickSort(points ScoreList) ScoreList {
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

	leftSide = score.QuickSort(leftSide)
	rightSide = score.QuickSort(rightSide)

	orderedArr := append(leftSide, middleScore)
	orderedArr = append(orderedArr, rightSide...)
	return orderedArr
}

func (score ScoreList) compareScores(scoreObj1, scoreObj2 Score) int {
	score1 := scoreObj1.Letter + scoreObj1.Position
	score2 := scoreObj2.Letter + scoreObj2.Position

	if score1 > score2 {
		return 1
	}

	if score1 < score2 {
		return -1
	}

	return 0
}

func (score ScoreList) GetTop(number int) ScoreList {
	if len(score) <= number {
		return score
	}

	return score[:number]
}
