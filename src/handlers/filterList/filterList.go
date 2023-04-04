package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matheus-osorio/go-term-solver/src/controllers"
	"github.com/matheus-osorio/go-term-solver/src/entry"
	"github.com/matheus-osorio/go-term-solver/src/response"
)

/*
 * This function is a lambda endpoint. The expected behavior for this function is to filter
 * a set of words using the rules of the wordle game.
 * It receives 3 main arguments:
 * Header: word-size = The size of the words you are sending. Wrong size words will be cut out
 * Body: wordList = The set of words that may be filtered
 * Rules: A set of rules. The rules must contain a letter and a rule. The rule must be one of three:
 *  - Green: Means that the letter exists in that spot
 *  - Yellow: Means that the letter exists in somewhere other than the spot
 *  - Grey: Means that the letter is not in the word
 * The function will parse every word and send you only the results that match the rules
 */
func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	wordSizeHeader, ok := event.Headers["word-size"]

	if !ok {
		return events.APIGatewayProxyResponse{Body: "Header word-size is required!", StatusCode: http.StatusBadRequest}, nil
	}

	wordSize, err := strconv.Atoi(wordSizeHeader)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Header word-size should be an integer!", StatusCode: http.StatusBadRequest}, nil
	}

	var eventBody entry.FilterListBodyObject
	err = json.Unmarshal([]byte(event.Body), &eventBody)
	if err != nil {
		responseBody := fmt.Sprintf("Body could not be parsed! Error: %v", err)
		return events.APIGatewayProxyResponse{Body: responseBody, StatusCode: http.StatusBadRequest}, nil
	}

	solver := controllers.ControllerFactory(wordSize, "")
	wordList, removedList := solver.FilterWords(eventBody)
	res := response.CreateResponse(wordList, removedList)
	bodyResponse, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{Body: string(bodyResponse), StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(handler)
}
