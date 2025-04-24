package runnerreg

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetInstallationID(t *testing.T) {
	// Set up your mock client
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Simulate a response
			r := io.NopCloser(bytes.NewReader([]byte(`[{"id": 123}]`)))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	// Test the function
	headers := http.Header{}
	id, err := GetInstallationID(mockClient, headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != 123 {
		t.Errorf("Expected ID 123, got %d", id)
	}
}

func TestGetInstallationToken(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Simulate a response for the installation token request
			r := io.NopCloser(bytes.NewReader([]byte(`{"token": "test-installation-token"}`)))
			return &http.Response{
				StatusCode: 201,
				Body:       r,
			}, nil
		},
	}

	// Test the function
	headers := http.Header{}
	token, err := GetInstallationToken(mockClient, 123, headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if token != "test-installation-token" {
		t.Errorf("Expected token 'test-installation-token', got '%s'", token)
	}
}

func TestGetRunnerToken(t *testing.T) {
	mockClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Simulate a response for the runner token request
			r := io.NopCloser(bytes.NewReader([]byte(`{"token": "test-runner-token"}`)))
			return &http.Response{
				StatusCode: 201,
				Body:       r,
			}, nil
		},
	}

	// Test the function
	installationToken := "test-installation-token"
	orgName := "test-org"
	token, err := GetRunnerToken(mockClient, installationToken, orgName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if token != "test-runner-token" {
		t.Errorf("Expected token 'test-runner-token', got '%s'", token)
	}
}
