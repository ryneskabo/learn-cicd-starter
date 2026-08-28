package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKeyShouldWork(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey fkjaldjf876564hjdfsh")
	apiKey, err := GetAPIKey(header);
	if err != nil {
		t.Fatalf("GetAPIKey function did not parse")
	}
	if apiKey != "fkjaldjf876564hjdfsh" {
		t.Fatalf("Did not get correct field or didn't get full key")
	}
}

func TestGetAPIKeyShouldNotWork(t *testing.T) {
	header := http.Header{}
	apiKey, err := GetAPIKey(header);
	if err == nil {
		t.Fatalf("Did not return error like it should")
	}
	if apiKey != "" {
		t.Fatalf("Parsed key that isn't there")
	}
}

func TestGetAPIKeyShouldNotWorkInvalidFormat(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey")
	apiKey, err := GetAPIKey(header);
	if err == nil {
		t.Fatalf("Did not return error like it should")
	}
	if apiKey != "" {
		t.Fatalf("Parsed key that isn't there")
	}
}

