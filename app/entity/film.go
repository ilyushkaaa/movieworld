package entity

import "time"

type Film struct {
	ID            uint64
	Name          string
	Description   string
	Duration      uint16
	MinAge        uint8
	Country       string
	ProducerName  string
	Actors        []*Actor
	Genres        []*Genre
	DateOfRelease time.Time
	Reviews       []*Review
}
