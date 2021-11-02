package model

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
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}

type ReviewEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Review
}
