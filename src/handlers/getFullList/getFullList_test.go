package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/matheus-osorio/go-term-solver/src/controllers/score"
)

func Test_ShouldGetFullListOfWords(t *testing.T) {
	response, _ := handler(events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"word-size": "5",
		},
	})
	responseEncoded := response.Body
	bodyResponse := []score.ScoreList{}
	json.Unmarshal([]byte(responseEncoded), &bodyResponse)

	if len(bodyResponse[0]) != 9980 {
		t.Errorf("Error! Expected answer to be of 9980 words. Got: %d", len(bodyResponse[0]))
	}
}
