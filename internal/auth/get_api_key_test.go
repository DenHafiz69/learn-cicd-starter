package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	h := http.Header{}

	api, err := GetAPIKey(h)
	if api != "" && err != errors.New("no authorization header included") {
		t.Error("Empty auth produce API and/or not provide the correct error.")
	}

	// Test for <2 value of auth
	h.Add("Authorization", "Test")
	api, err = GetAPIKey(h)
	if api != "" && err != errors.New("malformed authorization header") {
		t.Error("Malform data produce non-empty API key or wrong error")
	}

	// Test for first value != ApiKey
	h.Add("Authorization", "NotApiKey TEst")
	api, err = GetAPIKey(h)
	if api != "" && err != errors.New("malformed authorization header") {
		t.Error("Malform data produce non-empty API key or wrong error")
	}

}
