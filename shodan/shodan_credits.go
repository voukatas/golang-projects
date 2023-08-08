package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"log"
)

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

type APIInfo struct {
	QueryCredits int  `json:"query_credits"`
	ScanCredits  int  `json:"scan_credits"`
	Telnet       bool `json:"telnet"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (client *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, client.apiKey))
	if err != nil {
		return nil, errors.New("failed to fetch API information")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status: %d", res.StatusCode)
	}

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, errors.New("failed to decode API response")
	}

	return &ret, nil
}

func main() {
	apiKey := os.Getenv("SHODAN_API_KEY")

	client := NewClient(apiKey)
	apiInfo, err := client.APIInfo()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("QueryCreds:%v ScanCreds:%v\n", apiInfo.QueryCredits, apiInfo.ScanCredits)
}
