package main

import (
	"fmt"
	"log"

	"gopkg.in/tdage.v1"
)

func main() {
	pool := tdage.NewAgeData()

	userId := int64(1027242622)
	r, err := pool.GetDateAsJSON(userId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s %s", userId, r.Status, fmt.Sprintf("%02d/%d", r.Date.Month(), r.Date.Year()))
}
