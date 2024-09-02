package tdage

import (
	"testing"
)

func TestGetDateAsJSON(t *testing.T) {
	pool := NewAgeData()

	userId := int64(1027242622)
	r, err := pool.GetDateAsJSON(userId)
	if err != nil {
		t.Fatalf("Error getting date: %v", err)
	}

	if r.Status == "" {
		t.Error("Expected non-empty status")
	}

	if r.Date.IsZero() {
		t.Error("Expected a valid date")
	}
}
