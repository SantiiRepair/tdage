package tdage

import (
	"testing"
)

func TestGetDateAsJSON(t *testing.T) {
	pool := NewPool()

	userId := int64(5142525308)
	result := pool.GetDate(userId)

	if result.Status == "" {
		t.Error("Expected non-empty status")
	}

	if result.Date.IsZero() {
		t.Error("Expected a valid date")
	}
}
