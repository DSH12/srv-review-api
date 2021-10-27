package producer

import (
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/ozonmp/srv-review-api/internal/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-review-api/internal/mocks"
)

func TestStart(t *testing.T) {
	testEvents := []model.ReviewEvent{
		{
			ID:     1,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: nil,
		},
	}

	t.Run("cleanears", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)
		events := make(chan model.ReviewEvent)
		prod := Producer(NewKafkaProducer(2, sender, events, workerpool.New(1), repo))

		sender.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()
		repo.EXPECT().Remove(gomock.Eq([]uint64{testEvents[0].ID})).AnyTimes()

		prod.Start()

		for _, e := range testEvents {
			events <- e
		}
		prod.Close()
	})

	t.Run("updaters", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)
		events := make(chan model.ReviewEvent)
		prod := Producer(NewKafkaProducer(2, sender, events, workerpool.New(1), repo))
		sender.EXPECT().Send(gomock.Any()).Return(errors.New("Test error")).AnyTimes()
		repo.EXPECT().Unlock(gomock.Eq([]uint64{testEvents[0].ID})).AnyTimes()

		prod.Start()

		for _, e := range testEvents {
			events <- e
		}
		prod.Close()
	})

}
