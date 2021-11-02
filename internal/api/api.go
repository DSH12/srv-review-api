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

func (o *reviewAPI) DescribeReviewV1(
	ctx context.Context,
	req *pb.DescribeReviewV1Request,
) (*pb.DescribeReviewV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeReviewV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	review, err := o.repo.DescribeReview(ctx, req.ReviewId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeReviewV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if review == nil {
		log.Debug().Uint64("reviewId", req.ReviewId).Msg("review not found")
		totalReviewNotFound.Inc()

		return nil, status.Error(codes.NotFound, "review not found")
	}

	log.Debug().Msg("DescribeReviewV1 - success")

	return &pb.DescribeReviewV1Response{
		Value: &pb.Review{
			Id:  review.ID,
			Foo: review.Foo,
		},
	}, nil
}
