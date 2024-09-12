package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to hold the trivia response
type TriviaResponse struct {
	Results []struct {
		Question string `json:"question"`
		Answer   string `json:"correct_answer"`
	} `json:"results"`
}

// FetchTrivia fetches a trivia question from Open Trivia Database API
func FetchTrivia() (string, string, error) {
	// Make GET request to fetch trivia question
	resp, err := http.Get("https://opentdb.com/api.php?amount=1&type=multiple")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var trivia TriviaResponse
	if err := json.NewDecoder(resp.Body).Decode(&trivia); err != nil {
		return "", "", err
	}

	// Return the first question and answer from the API response
	if len(trivia.Results) > 0 {
		question := trivia.Results[0].Question
		answer := trivia.Results[0].Answer
		return question, answer, nil
	}
	return "", "", fmt.Errorf("no trivia found")
}
