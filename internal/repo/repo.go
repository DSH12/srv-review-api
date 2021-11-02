package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/srv-review-api/internal/model"
)

// Repo is DAO for Review
type Repo interface {
	DescribeReview(ctx context.Context, reviewID uint64) (*model.Review, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeReview(ctx context.Context, reviewID uint64) (*model.Review, error) {
	return nil, nil
}
