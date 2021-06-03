/*

This package contains code to take in a location string and
return news for that location

*/
package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nimitpatel26/tenki/go/news"
)

func getWeather(location string) string {
	data, _ := json.Marshal(news.GetNews(location))
	return string(data)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       getWeather(request.QueryStringParameters["location"]),
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
