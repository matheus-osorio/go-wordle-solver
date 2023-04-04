package response

import "github.com/matheus-osorio/go-term-solver/src/controllers/score"

type Response struct {
	AvailableWords []score.ScoreList `json:"availableWords"`
	RemovedWords   []score.ScoreList `json:"removedWords"`
}

func CreateResponse(availableWords, removedWords []score.ScoreList) Response {
	return Response{
		AvailableWords: availableWords,
		RemovedWords:   removedWords,
	}
}
