package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestShouldGetList(t *testing.T) {
	event := events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"word-size": "5",
			"language":  "pt-br",
		},
	}

	t.Setenv("IS_OFFLINE", "true")

	_, err := handler(event)

	if err != nil {
		fmt.Printf("%v", err)
		t.Errorf("Error when calling function")
	}

}
