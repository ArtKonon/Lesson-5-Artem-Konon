package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const aiReviewPrompt = "Review this Go code for missing or incomplete error handling. List every location where an error is ignored, mishandled, or under-wrapped, and propose a fix for each using best practices like fmt.Errorf with %w"

type badUser struct {
	Name string `json:"name"`
}

func badLoadUser(path string) (badUser, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		_ = err
	}

	if len(data) == 0 {
		panic("validation failed: user file is empty")
	}

	var u badUser
	if err := json.Unmarshal(data, &u); err != nil {
		return badUser{}, fmt.Errorf("decode user: %v", err)
	}

	return u, nil
}

func badSaveUser(path string, u badUser) error {
	data, err := json.Marshal(u)
	if err != nil {
		_ = err
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write user: %v", err)
	}

	return nil
}
