package strava

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type StravaClient struct {
	http.Client
	target *url.URL
}

func NewStravaClient(target *url.URL) *StravaClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Failed to create cookie jar: %s", err)
	}
	return &StravaClient{http.Client{Jar: jar}, target}
}

func (client *StravaClient) send(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return nil, fmt.Errorf("Request '%s' failed with status code %s", req.URL, resp.Status)
	}

	return resp, nil
}

func newStravaLoginRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://www.strava.com/login", nil)
	if err != nil {
		return nil, err
	}
	setBrowserCookies(req)

	return req, nil
}

func newStravaSessionRequest(email string, password string, token string) (*http.Request, error) {
	data := url.Values{}
	data.Add("utf8", "\u2713")
	data.Add("authenticity_token", token)
	data.Add("plan", "")
	data.Add("email", email)
	data.Add("password", password)
	data.Add("remember_me", "on")

	req, err := http.NewRequest("POST", "https://www.strava.com/session", strings.NewReader(data.Encode()))
	if err != nil {
		setBrowserCookies(req)
	}

	return req, err
}

func extractAuthenticityToken(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read body: %w", err)
	}

	expr, _ := regexp.Compile("name=\"authenticity_token\".*value=\"(.+?)\"")
	matches := expr.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", errors.New("Token not found")
	}

	return matches[1], nil
}

func (client *StravaClient) Authenticate(email string, password string) error {
	res, err := client.send(newStravaLoginRequest())
	if err != nil {
		return fmt.Errorf("Could not send login request: %w", err)
	}
	token, err := extractAuthenticityToken(res)
	if err != nil {
		return fmt.Errorf("Could not get authenticity token for login form: %w", err)
	}
	if _, err = client.send(newStravaSessionRequest(email, password, token)); err != nil {
		return fmt.Errorf("Could not send session request: %w", err)
	}
	auth, _ := url.JoinPath(client.target.String(), "auth")
	if _, err := client.send(http.NewRequest("GET", auth, nil)); err != nil {
		return fmt.Errorf("Could not send auth request: %w", err)
	}
	return nil
}

func (client *StravaClient) GetCloudFrontCookies() map[string]string {
	cookies := map[string]string{}
	target, _ := url.Parse("https://www.strava.com")
	for _, cookie := range client.Jar.Cookies(target) {
		if strings.HasPrefix(cookie.Name, "CloudFront") {
			cookies[cookie.Name] = cookie.Value
		}
	}
	return cookies
}

func (client *StravaClient) AddCookies(req *http.Request) {
	for _, c := range client.Jar.Cookies(req.URL) {
		req.AddCookie(c)
	}
}

func (client *StravaClient) GetTarget() *url.URL {
	return client.target
}

func setBrowserCookies(req *http.Request) {
	if req == nil {
		return
	}

	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rn.Intn(len(userAgents))
	randomUserAgent := userAgents[randomIndex]

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Referer", "https://www.strava.com/login")
	req.Header.Set("Sec-Ch-Ua", randomUserAgent.SecChUa)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", randomUserAgent.UserAgent)
}
