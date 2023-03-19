package score

import "testing"

func Test_ShouldCompareScores(t *testing.T) {
	score := ScoreList{}
	// arrange
	t.Parallel()
	score1 := Score{
		Points: 1.0,
	}

	score2 := Score{
		Points: 2.0,
	}

	// act
	result1 := score.compareScores(score1, score2)
	score1.Points = 2.0

	result2 := score.compareScores(score1, score2)
	score1.Points = 3.0

	result3 := score.compareScores(score1, score2)

	// assert
	if result1 != -1 {
		t.Errorf("Expected -1 result, got %d", result1)
	}

	if result2 != 0 {
		t.Errorf("Expected 0 result, got %d", result2)
	}

	if result3 != 1 {
		t.Errorf("Expected 1 result, got %d", result3)
	}
}

func Test_ShouldSortByScore(t *testing.T) {
	t.Parallel()

	// arrange
	scoreList := ScoreList{
		{
			Word:   "word1",
			Points: 0.0,
		},
		{
			Word:   "word2",
			Points: 5.0,
		},
		{
			Word:   "word3",
			Points: 10.0,
		},
	}

	// act
	result := scoreList.SortByScore()

	// assert
	if result[0].Word != "word3" {
		t.Errorf("Score was not sorted properly. %v", result)
	}
}
