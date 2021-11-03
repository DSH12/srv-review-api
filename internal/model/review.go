package model

import "time"

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type Review struct {
	ID       uint64    `db:"id"`
	AuthorID uint64    `db:"author_id"`
	Text     string    `db:"text"`
	Score    uint32    `db:"score"`
	Created  time.Time `db:"created"`
}

type ReviewEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Review
}
