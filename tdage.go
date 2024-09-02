package tdage

import (
	"math"
	"sort"
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
//
//	A pointer to a new instance of AgeData, with the minID and maxID
//	fields set according to the IDs found in the dataset.
func NewPool() *AgeData {
	a := &AgeData{}

	for k := range ages {
		if a.minID == 0 || k < a.minID {
			a.minID = k
		}
		if a.maxID == 0 || k > a.maxID {
			a.maxID = k
		}
	}

	return a
}

// Calculates the creation date based on the provided integer value.
//
// It returns a tuple containing the result of the calculation and the date of creation as a time.Time object.
//
// Parameters:
//
//	v: An integer representing the value used to determine the creation date.
//
// Returns:
//
//	A tuple containing the result of the calculation and the date of creation as a time.Time object.
func (a *AgeData) GetAsDatetime(v int64) (int, time.Time) {
	if v < a.minID {
		return -1, time.Unix(ages[a.minID]/1000, 0)
	} else if v > a.maxID {
		return 1, time.Unix(ages[a.maxID]/1000, 0)
	}

	keys := make([]int64, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	lowerID := a.minID
	for _, k := range keys {
		if v <= k {
			lage := float64(ages[lowerID]) / 1000
			uage := float64(ages[k]) / 1000
			vRatio := float64(v-lowerID) / float64(k-lowerID)
			midDate := math.Floor((vRatio * (uage - lage)) + lage)
			return 0, time.Unix(int64(midDate), 0)
		}
		lowerID = k
	}

	return 0, time.Time{}
}

// GetDate calculates the creation date based on the provided integer value.
//
// It returns a Result containing the status of the date comparison and the corresponding date.
//
// Parameters:
//
//	v: An integer representing the value used to determine the creation date.
//
// Returns:
//
//	A Result struct containing the status of the comparison ("older_than", "newer_than", or "aprox")
//	and the date as a time.Time object, along with an error if applicable.
func (a *AgeData) GetDate(v int64) Result {
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

	return Result{Status: status, Date: d}
}
