package tdage

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"
)

type AgeData struct {
	data  map[string]int64
	minID int64
	maxID int64
}

type Result struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func NewAgeData() *AgeData {
	return &AgeData{
		data: make(map[string]int64),
	}
}

func (a *AgeData) LoadAges(filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileData, &a.data); err != nil {
		return err
	}

	for k := range a.data {
		id := parseID(k)
		if a.minID == 0 || id < a.minID {
			a.minID = id
		}
		if a.maxID == 0 || id > a.maxID {
			a.maxID = id
		}
	}
	return nil
}

func parseID(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}

func (a *AgeData) GetDateAsDatetime(v int64) (int, time.Time) {
	if v < a.minID {
		return -1, time.Unix(a.data[fmt.Sprintf("%d", a.minID)]/1000, 0)
	} else if v > a.maxID {
		return 1, time.Unix(a.data[fmt.Sprintf("%d", a.maxID)]/1000, 0)
	}

	lowerID := a.minID
	for k, v := range a.data {
		uid := parseID(k)
		if v <= uid {
			lage := float64(a.data[fmt.Sprintf("%d", lowerID)]) / 1000
			uage := float64(v) / 1000
			idRatio := float64(v-lowerID) / float64(uid-lowerID)
			midDate := math.Floor((idRatio * (uage - lage)) + lage)
			return 0, time.Unix(int64(midDate), 0)
		}
		lowerID = uid
	}

	return 0, time.Time{}
}

func (a *AgeData) GetDateAsJSON(v int64) (Result, error) {
	s, d := a.GetDateAsDatetime(v)
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
