package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	. "main/lib"
)

type MyEvent struct {
	Url string `json:"url"`
}

type MyResponse struct {
	Url        string           `json:"Url:"`
	Resolution map[string]Match `json:"Resolution:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Url: event.Url, Resolution: Resolve(event.Url)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
