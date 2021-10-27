package sender

import (
	"github.com/ozonmp/srv-review-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.ReviewEvent) error
}
