package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// AI review prompt used to identify and fix errors in the bad example.
// "Review this Go code for missing or incomplete error handling. List every location where an error is ignored, mishandled, or under-wrapped, and propose a fix for each using best practices like fmt.Errorf with %w"

type fixedUser struct {
	Name string `json:"name"`
}

func fixedLoadUser(path string) (fixedUser, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return fixedUser{}, fmt.Errorf("read user file %q: %w", path, err)
	}

	if len(data) == 0 {
		return fixedUser{}, fmt.Errorf("read user file %q: %w", path, os.ErrInvalid)
	}

	var u fixedUser
	if err := json.Unmarshal(data, &u); err != nil {
		return fixedUser{}, fmt.Errorf("decode user data from %q: %w", path, err)
	}

	return u, nil
}

func fixedSaveUser(path string, u fixedUser) error {
	data, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("marshal user: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write user file %q: %w", path, err)
	}

	return nil
}
