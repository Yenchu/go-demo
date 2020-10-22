package model

import (
	"time"
)

type Data struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Location  Location  `json:"location,omitempty"`
	DateAdded time.Time `db:"date_added" json:"dateAdded,omitempty"`
}

type Location struct {
	Lat float32 `db:"latitude" json:"lat,omitempty"`
	Lng float32 `db:"longitude" json:"lng,omitempty"`
}
