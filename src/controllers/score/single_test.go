package score

import (
	"math"
	"testing"
)

func Test_ShouldGetWordScore(t *testing.T) {
	// arrange
	t.Parallel()
	wordList := []string{"aaaaa", "bbbbb", "ccccc"}

	scoreSystem := SingleScoreSystem{
		WordList:   wordList,
		WordSize:   5,
		TotalWords: 3,
	}

	scoreSystem.createScoreTable()
	type Test struct {
		EntryObject string
		Score       float64
	}
	testEntries := []Test{
		{
			EntryObject: "aaaaa",
			Score:       5 * (2.0*(1.0/3.0) + 0*(0/3.0) + 1.0*(2.0/3.0)),
		},
	}

	scoreSystem.createScoreTable()

	for _, expectedWords := range testEntries {
		// act
		score := scoreSystem.getWordScore(expectedWords.EntryObject)

		// assert
		if math.Round(expectedWords.Score) != math.Round(score) {
			t.Errorf("ERROR: Expected %s score to be: %f. Got: %f", expectedWords.EntryObject, expectedWords.Score, score)
		}

	}
}

func Test_ShouldCreateScoreTable(t *testing.T) {
	t.Parallel()
	// arrange
	wordList := []string{"aaaaa", "abbbb", "acccc"}

	scoreSystem := SingleScoreSystem{
		WordList: wordList,
		WordSize: 5,
	}

	// act
	scoreSystem.createScoreTable()

	// assert
	if score := scoreSystem.PositionPoints[0]['a']; score != 3 {
		t.Errorf("Expected a score of 3, found %f", score)
	}

	if score := scoreSystem.PositionPoints[0]['b']; score != 0 {
		t.Errorf("Expected a score of 0, found %f", score)
	}

	if score := scoreSystem.PositionPoints[4]['c']; score != 1 {
		t.Errorf("Expected a score of 1, found %f", score)
	}
}

func Test_ShouldSetScores(t *testing.T) {
	score := SingleScoreSystem{
		WordList: []string{"aaaaa", "bbbbb", "ccccc"},
		WordSize: 5,
	}

	score.createScoreTable()
	score.setScores()

	if len(score.FilteredPoints) != 3 {
		t.Errorf("Expected score list to be the same size as word list. Got %d", len(score.FilteredPoints))
	}
}
