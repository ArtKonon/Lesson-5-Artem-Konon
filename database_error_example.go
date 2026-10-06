package main

import (
	"errors"
	"fmt"
)

// DatabaseError stores the failing SQL query and the original error.
type DatabaseError struct {
	Query       string
	OriginalErr error
}

func (e *DatabaseError) Error() string {
	if e == nil {
		return "database error"
	}
	if e.OriginalErr == nil {
		return fmt.Sprintf("database error while executing query: %s", e.Query)
	}
	return fmt.Sprintf("database error while executing query %q: %v", e.Query, e.OriginalErr)
}

func FetchUser(id int) error {
	users := map[int]string{42: "alice"}

	if _, ok := users[id]; !ok {
		err := fmt.Errorf("user %d not found", id)
		return fmt.Errorf("SELECT id, name FROM users WHERE id = %d: %w", id, &DatabaseError{
			Query:       "SELECT id, name FROM users WHERE id = ?",
			OriginalErr: err,
		})
	}

	return nil
}

func main() {
	if err := FetchUser(404); err != nil {
		var dbErr *DatabaseError
		if errors.As(err, &dbErr) {
			fmt.Printf("database query failed: %s\n", dbErr.Query)
			fmt.Println(err)
			return
		}
		fmt.Println(err)
	}
}
