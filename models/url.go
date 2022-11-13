package models

import "time"

type Url struct {
	Url       string
	Id        uint32
	CreatedAt time.Time
	TTL       int
}
