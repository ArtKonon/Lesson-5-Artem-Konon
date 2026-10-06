package main

import (
	"errors"
	"testing"
)

func TestFetchUserReturnsWrappedDatabaseError(t *testing.T) {
	err := FetchUser(404)
	if err == nil {
		t.Fatal("expected an error")
	}

	var dbErr *DatabaseError
	if !errors.As(err, &dbErr) {
		t.Fatalf("expected DatabaseError via errors.As: %v", err)
	}

	if dbErr.Query == "" {
		t.Fatal("expected a non-empty query string")
	}

	if dbErr.OriginalErr == nil {
		t.Fatal("expected an original error to be attached")
	}
}
