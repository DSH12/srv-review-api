package producer

import (
	"sync"
	"time"

	"github.com/ozonmp/srv-review-api/internal/app/repo"
	"github.com/ozonmp/srv-review-api/internal/app/sender"
	"github.com/ozonmp/srv-review-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	events <-chan model.ReviewEvent

	repo repo.EventRepo

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan model.ReviewEvent,
	workerPool *workerpool.WorkerPool,
	repo repo.EventRepo,
) *producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
		repo:       repo,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							p.repo.Unlock([]uint64{event.ID})
						})
					} else {
						p.workerPool.Submit(func() {
							p.repo.Remove([]uint64{event.ID})
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
