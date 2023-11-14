package runnerreg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structs for parsing JSON responses
type Installation struct {
	ID int `json:"id"`
}

type Token struct {
	Token string `json:"token"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func GetInstallationID(client HttpClient, headers http.Header) (int, error) {
	url := "https://api.github.com/app/installations"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var installations []Installation
	err = json.NewDecoder(resp.Body).Decode(&installations)
	if err != nil {
		return 0, err
	}

	return installations[0].ID, nil
}

func GetInstallationToken(client HttpClient, installationID int, headers http.Header) (string, error) {
	url := fmt.Sprintf("https://api.github.com/app/installations/%d/access_tokens", installationID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}

func GetRunnerToken(client HttpClient, installationToken, orgName string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/actions/runners/registration-token", orgName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	headers := http.Header{
		"Authorization": []string{fmt.Sprintf("token %s", installationToken)},
		"Accept":        []string{"application/vnd.github.v3+json"},
	}
	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
