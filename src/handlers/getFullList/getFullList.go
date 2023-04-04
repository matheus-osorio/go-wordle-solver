package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matheus-osorio/go-term-solver/src/controllers"
	"github.com/matheus-osorio/go-term-solver/src/entry"
	"github.com/matheus-osorio/go-term-solver/src/response"
)

/*
 * This function is a lambda endpoint. The expected behavior of this function is
 * to receive a language and a word size and return the word values that match
 * The necessary parameters are:
 *  - Header: word-size(uint) = Size of the words
 *  - Header: language(string) = Language of the words
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

	language, ok := event.Headers["language"]

	if !ok {
		return events.APIGatewayProxyResponse{Body: "Header language is required!", StatusCode: http.StatusBadRequest}, nil
	}

	if !entry.AVAILABLE_LANGUAGES[language] {
		return events.APIGatewayProxyResponse{Body: "Language not accepted!", StatusCode: http.StatusBadRequest}, nil
	}

	solver := controllers.ControllerFactory(wordSize, language)
	wordList, removedList := solver.GetFullWordList()
	res := response.CreateResponse(wordList, removedList)
	bodyResponse, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{Body: string(bodyResponse), StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(handler)
}
