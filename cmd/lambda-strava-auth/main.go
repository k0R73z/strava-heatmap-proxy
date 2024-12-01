package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/k0R73z/strava-heatmap-proxy/internal/strava"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	email := request.QueryStringParameters["email"]
	password := request.QueryStringParameters["password"]
	providerURL := request.QueryStringParameters["providerURL"]
	if providerURL == "" {
		providerURL = "https://heatmap-external-a.strava.com/"
	}

	config, err := strava.NewConfig(email, password)
	if err != nil {
		return response(400, fmt.Sprintf("%s", err)), nil
	}

	target, err := url.Parse(providerURL)
	if err != nil {
		return response(400, fmt.Sprintf("Invalid providerURL=%s", err)), nil
	}

	client := strava.NewStravaClient(target)
	if err := client.Authenticate(config.Email, config.Password); err != nil {
		msg := fmt.Sprintf("Failed to authenticate client: %s", err)
		return response(500, msg), nil
	}

	cookies := client.GetCloudFrontCookies()

	jsonData, err := json.Marshal(cookies)
	if err != nil {
		return response(500, fmt.Sprintf("Can't marshall json output=%s", err)), nil
	}

	jsonString := string(jsonData)
	return response(200, jsonString), nil
}

func response(code int, message string) events.APIGatewayProxyResponse {
	if code == 500 {
		log.Printf("Error=%s\n", message)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       message,
	}
}

func main() {
	lambda.Start(handler)
}
