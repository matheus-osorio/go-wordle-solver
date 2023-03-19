package score

// Factory method for the Score package. It can be used to give the correct score for single and double modes
// TODO: Create the Double Score System
// TODO: Create the Quadruple Score System (maybe a non-single method may suffice)
func CreateScore(wordList [][]string, wordSize int) ScoreInterface {
	switch len(wordList) {
	case 1:
		return &SingleScoreSystem{
			WordList:   wordList[0],
			WordSize:   wordSize,
			TotalWords: len(wordList[0]),
		}

	default:
		panic("Not implemented other types!!!")
	}
}
