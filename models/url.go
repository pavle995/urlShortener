package models

import "time"

type Url struct {
	Url       string    `json:"fullUrl"`
	Id        uint32    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	TTL       int       `json:"ttl"`
}
