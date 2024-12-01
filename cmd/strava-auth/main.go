package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/k0R73z/strava-heatmap-proxy/internal/strava"
)

var providerURL, email, password string

func main() {
	flag.StringVar(&providerURL, "providerURL", "https://heatmap-external-a.strava.com/", "Heatmap provider URL")
	flag.StringVar(&email, "email", "", "Login email")
	flag.StringVar(&password, "password", "", "Login password")

	flag.Parse()

	config, err := strava.NewConfig(email, password)
	if err != nil {
		log.Fatal(err)
	}

	target, err := url.Parse(providerURL)
	if err != nil {
		log.Fatalf("Invalid providerURL=%s", err)
	}

	client := strava.NewStravaClient(target)
	if err := client.Authenticate(config.Email, config.Password); err != nil {
		log.Fatalf("Failed to authenticate client: %s", err)
	}

	cookies := client.GetCloudFrontCookies()

	fmt.Fprintf(os.Stdout,
		"CloudFront-Key-Pair-Id=%s CloudFront-Policy=%s CloudFront-Signature=%s\n",
		cookies["CloudFront-Key-Pair-Id"],
		cookies["CloudFront-Policy"],
		cookies["CloudFront-Signature"])
}
