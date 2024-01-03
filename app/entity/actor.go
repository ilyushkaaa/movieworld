package entity

import "time"

type Actor struct {
	ID          uint64
	Name        string
	Surname     string
	Nationality string
	Birthday    time.Time
}
