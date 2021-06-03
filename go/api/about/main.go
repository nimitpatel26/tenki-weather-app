/*

This package contains code to
return about data for the weather application

*/
package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nimitpatel26/tenki/go/fauna"
)

func getAboutData() string {
	data, _ := json.Marshal(fauna.GetAboutData())
	return string(data)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       getAboutData(),
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
