package usecase

import (
	"context"
	"io"
	"port-processor/internal/entity"
)

type (
	Port interface {
		Process(ctx context.Context, reader io.Reader, workerCount int) error
	}

	// PortRepo is the repository interface for Port
	PortRepo interface {
		Save(port *entity.Port) error
	}
)
