package waifupicsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.waifu.pics/"

func get(endpoint string) (*http.Response, error) {
	url := baseURL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	return resp, nil
}

func post(endpoint string, payload any) (*http.Response, error) {
	url := baseURL + endpoint

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("POST %s: %w", url, err)
	}
	return resp, nil
}
