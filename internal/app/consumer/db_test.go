package consumer

import (
	//"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-review-api/internal/mocks"
	"github.com/ozonmp/srv-review-api/internal/model"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)

	testEvents := []model.ReviewEvent{{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: nil,
	},
	}
	events := make(chan model.ReviewEvent, 512)
	con := Consumer(NewDbConsumer(2, 2, 5*time.Second, repo, events))

	repo.EXPECT().Lock(gomock.Any()).Return(testEvents, nil).AnyTimes()

	con.Start()

	checkEvent := <-events
	if checkEvent != testEvents[0] {
		t.Fail()
	}

	con.Close()

}
