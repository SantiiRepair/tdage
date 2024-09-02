package tdage

import (
	"fmt"
	"math"
	"time"
)

type AgeData struct {
	minID int64
	maxID int64
}

type Result struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

// NewPool creates and returns a new instance of AgeData.
//
// This function initializes a new AgeData object and calculates the
// minimum and maximum ID values from a dataset called `ages`.
//
// Returns:
//   A pointer to a new instance of AgeData, with the minID and maxID
//   fields set according to the IDs found in the dataset.
func NewPool() *AgeData {
	a := &AgeData{}

	for k := range ages {
		id := parseID(k)
		if a.minID == 0 || id < a.minID {
			a.minID = id
		}
		if a.maxID == 0 || id > a.maxID {
			a.maxID = id
		}
	}

	return a
}

func parseID(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}

// Calculates the creation date based on the provided integer value.
//
// It returns a tuple containing the result of the calculation and the date of creation as a time.Time object.
//
// Parameters:
//   v: An integer representing the value used to determine the creation date.
//
// Returns:
//   A tuple containing the result of the calculation and the date of creation as a time.Time object.
func (a *AgeData) GetAsDatetime(v int64) (int, time.Time) {
	if v < a.minID {
		return -1, time.Unix(ages[fmt.Sprintf("%d", a.minID)]/1000, 0)
	} else if v > a.maxID {
		return 1, time.Unix(ages[fmt.Sprintf("%d", a.maxID)]/1000, 0)
	}

	lowerID := a.minID
	for x, y := range ages {
		uid := parseID(x)
		if v <= uid {
			lage := float64(ages[fmt.Sprintf("%d", lowerID)]) / 1000
			uage := float64(y) / 1000
			idRatio := float64(v-lowerID) / float64(uid-lowerID)
			midDate := math.Floor((idRatio * (uage - lage)) + lage)
			return 0, time.Unix(int64(midDate), 0)
		}

		lowerID = uid
	}

	return 0, time.Time{}
}

// GetDate calculates the creation date based on the provided integer value.
//
// It returns a Result containing the status of the date comparison and the corresponding date.
//
// Parameters:
//   v: An integer representing the value used to determine the creation date.
//
// Returns:
//   A Result struct containing the status of the comparison ("older_than", "newer_than", or "aprox")
//   and the date as a time.Time object, along with an error if applicable.
func (a *AgeData) GetDate(v int64) (Result, error) {
	s, d := a.GetAsDatetime(v)
	var status string
	switch s {
	case -1:
		status = "older_than"
	case 1:
		status = "newer_than"
	default:
		status = "aprox"
	}

	return Result{Status: status, Date: d}, nil
}
