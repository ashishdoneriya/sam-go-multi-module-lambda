package main

import (
	"fmt"
	"my-multi-module-project/common" // Import the common module locally

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, service2 !\n" + common.SharedFunction()
	} else {
		greeting = fmt.Sprintf("Hello, service2 %s!\n"+common.SharedFunction(), sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
