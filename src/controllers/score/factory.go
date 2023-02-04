package score

func CreateFactory(wordList [][]string, wordSize int) ScoreInterface {
	switch len(wordList) {
	case 1:
		return &SingleScoreSystem{
			WordList: wordList[0],
			WordSize: wordSize,
		}

	default:
		panic("Not implemented other types!!!")
	}
}
