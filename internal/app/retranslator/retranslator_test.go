package retranslator

import (
	"errors"
	"github.com/ozonmp/srv-review-api/internal/model"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-review-api/internal/mocks"
)

func TestStart(t *testing.T) {

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           nil,
		Sender:         nil,
	}

	testEvents := []model.ReviewEvent{
		{
			ID:     1,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: nil,
		},
	}

	t.Parallel()

	t.Run("cleanears", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		repo.EXPECT().Lock(gomock.Any()).Return(testEvents, nil).AnyTimes()
		sender.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()
		repo.EXPECT().Remove(gomock.Eq([]uint64{testEvents[0].ID})).AnyTimes()

		cfg.Repo = repo
		cfg.Sender = sender

		retranslator := Retranslator(NewRetranslator(cfg))
		retranslator.Start()
		retranslator.Close()
	})

	t.Run("updaters", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		repo.EXPECT().Lock(gomock.Any()).Return(testEvents, nil).AnyTimes()
		sender.EXPECT().Send(gomock.Any()).Return(errors.New("Test error")).AnyTimes()
		repo.EXPECT().Unlock(gomock.Eq([]uint64{testEvents[0].ID})).AnyTimes()

		cfg.Repo = repo
		cfg.Sender = sender

		retranslator := Retranslator(NewRetranslator(cfg))
		retranslator.Start()
		retranslator.Close()
	})

}
