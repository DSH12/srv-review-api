package api

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/srv-review-api/internal/repo"

	pb "github.com/ozonmp/srv-review-api/pkg/srv-review-api"
)

var (
	totalReviewNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "srv_review_api_review_not_found_total",
		Help: "Total number of reviews that were not found",
	})
)

type reviewAPI struct {
	pb.UnimplementedSrvReviewApiServiceServer
	repo repo.Repo
}

// NewReviewAPI returns api of srv-review-api service
func NewReviewAPI(r repo.Repo) pb.SrvReviewApiServiceServer {
	return &reviewAPI{repo: r}
}

func (o *reviewAPI) DescribeReviewV1(ctx context.Context, req *pb.DescribeReviewV1Request,
) (*pb.DescribeReviewV1Response, error) {
	log.Debug().Msg("Describe: " + req.String())
	return nil, status.Error(codes.Unimplemented, "Describe not implemented")
}

func (o *reviewAPI) CreateReviewV1(ctx context.Context, req *pb.CreateReviewV1Request,
) (*pb.CreateReviewV1Response, error) {
	log.Debug().Msg("Create: " + req.String())
	return nil, status.Error(codes.Unimplemented, "Create not implemented")
}

func (o *reviewAPI) ListReviewV1(ctx context.Context, req *pb.ListReviewV1Request,
) (*pb.ListReviewV1Response, error) {
	log.Debug().Msg("List: " + req.String())
	return nil, status.Error(codes.Unimplemented, "List not implemented")

}
func (o *reviewAPI) RemoveReviewV1(ctx context.Context, req *pb.RemoveReviewV1Request,
) (*pb.RemoveReviewV1Response, error) {
	log.Debug().Msg("Remove: " + req.String())
	return nil, status.Error(codes.Unimplemented, "Remove not implemented")
}
