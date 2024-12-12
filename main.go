package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/helmfile/vals"
	runnerreg "github.com/remarkable/gh-runner-regger/pkg/github"
)

func getJWT(pemBytes []byte, appID string) (string, error) {
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemBytes)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"iss": appID,
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getSecretVal(pemPath string) ([]byte, error) {
	// get the secret value
	runtime, err := vals.New(vals.Options{
		LogOutput:             os.Stderr,
		CacheSize:             100,
		ExcludeSecret:         true,
		FailOnMissingKeyInMap: true,
	})
	if err != nil {
		return nil, err
	}

	signString, err := runtime.Get(pemPath)
	if err != nil {
		return nil, err
	}
	return []byte(signString), nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <path/to/private-key.pem|ref+$vals-expression> <app-id> <org>\n", os.Args[0])
		os.Exit(1)
	}

	pemPath := os.Args[1]
	appID := os.Args[2]
	org := os.Args[3]

	// if pemPath starts with the string "ref+", we should use vals to fetch the value
	var signBytes []byte
	var err error
	if strings.HasPrefix(pemPath, "ref+") {
		signBytes, err = getSecretVal(pemPath)
	} else {
		signBytes, err = os.ReadFile(pemPath)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tokenString, err := getJWT(signBytes, appID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	headers := http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", tokenString)},
		"Accept":        []string{"application/vnd.github.v3+json"},
	}
	client := &http.Client{}

	installationID, err := runnerreg.GetInstallationID(client, headers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	installationToken, err := runnerreg.GetInstallationToken(client, installationID, headers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	runnerToken, err := runnerreg.GetRunnerToken(client, installationToken, org)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(runnerToken)
}
