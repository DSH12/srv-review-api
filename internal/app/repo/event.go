package repo

import (
	"github.com/ozonmp/srv-review-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.ReviewEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.ReviewEvent) error
	Remove(eventIDs []uint64) error
}
