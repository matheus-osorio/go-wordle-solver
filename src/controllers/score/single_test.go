package score

import (
	"testing"
)

func Test_ShouldGetWordScore(t *testing.T) {
	// arrange
	t.Parallel()
	wordList := []string{"aaaaa", "bbbbb", "ccccc"}

	scoreSystem := SingleScoreSystem{
		WordList: wordList,
		WordSize: 5,
	}

	scoreSystem.createScoreTable()
	type Test struct {
		EntryObject   string
		LetterScore   float64
		PositionScore float64
	}
	testEntries := []Test{
		{
			EntryObject:   "affff",
			LetterScore:   4.0 / 16,
			PositionScore: 7,
		},
		{
			EntryObject:   "bkkkk",
			LetterScore:   4.0 / 16,
			PositionScore: 1,
		},
		{
			EntryObject:   "cpppp",
			LetterScore:   4.0 / 16,
			PositionScore: 1,
		},
		{
			EntryObject:   "aaaaa",
			LetterScore:   20.0 / 16,
			PositionScore: 5,
		},
		{
			EntryObject:   "adddd",
			LetterScore:   4.0 / 16,
			PositionScore: 1,
		},
	}

	for _, expectedWords := range testEntries {
		// act
		letterScore, positionScore := scoreSystem.getWordScore(expectedWords.EntryObject)

		// assert
		if expectedWords.PositionScore != positionScore {
			t.Errorf("ERROR: Expected %s position score to be: %f. Got: %f", expectedWords.EntryObject, expectedWords.PositionScore, positionScore)
		}

		if expectedWords.LetterScore != letterScore {
			t.Errorf("ERROR: Expected %s letter score to be: %f. Got: %f", expectedWords.EntryObject, expectedWords.LetterScore, letterScore)
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
	if score := scoreSystem.ScorePoints[0]['a']; score != 3 {
		t.Errorf("Expected a score of 3, found %f", score)
	}

	if score := scoreSystem.ScorePoints[0]['b']; score != 0 {
		t.Errorf("Expected a score of 0, found %f", score)
	}

	if score := scoreSystem.ScorePoints[4]['c']; score != 1 {
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

	if len(score.Points) != 3 {
		t.Errorf("Expected score list to be the same size as word list. Got %d", len(score.Points))
	}
}
