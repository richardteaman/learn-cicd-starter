package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if key != "test123" {
		t.Errorf("Expected key to be 'test123', got '%s'", key)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer sometoken")

	_, err := GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header' error, got %v", err)
	}
}
