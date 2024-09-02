package tdage

import (
	"testing"
)

func TestGetDateAsJSON(t *testing.T) {
	pool := NewPool()

	userId := int64(5142525308)
	r, err := pool.GetDate(userId)
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
