package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"port-processor/internal/entity"
	"sync"
)

// PortUseCase is a use case for port.
type PortUseCase struct {
	repo PortRepo
}

// NewPortUseCase creates a new PortUseCase.
func NewPortUseCase(r PortRepo) *PortUseCase {
	return &PortUseCase{
		repo: r,
	}
}

// to reduce memory allocations
var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() any {
			return new(entity.Port)
		},
	}
}

func (p *PortUseCase) Process(ctx context.Context, reader io.Reader, workerCount int) (err error) {

	dec := json.NewDecoder(reader)
	// read open bracket
	var t json.Token
	t, err = dec.Token()
	if err != nil {
		return
	}

	done := make(chan struct{})
	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, workerCount)
	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
	}()

	go func() {
		defer close(done)
		// while the array contains values
		for dec.More() {
			// unloc of port
			t, err = dec.Token()
			if err != nil {
				return
			}

			if !dec.More() {
				continue
			}

			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			port := pool.Get().(*entity.Port)
			err = dec.Decode(port)
			if err != nil {
				return
			}
			// set unloc
			port.Unloc = fmt.Sprintf("%v", t)

			// parallelize db operations
			go func(_port *entity.Port) {
				defer func() {
					pool.Put(_port)
					<-semaphoreChan
				}()
				// TODO: add error handling
				p.repo.Save(_port)
			}(port)
		}
	}()

	<-done
	return
}
