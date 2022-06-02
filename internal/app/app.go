package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"port-processor/internal/entity"
	"port-processor/internal/http"
	"port-processor/internal/usecase"
	"port-processor/internal/usecase/repo"
	"port-processor/pkg/config"
	"port-processor/pkg/db"
	"port-processor/pkg/httpserver"
	"syscall"
)

// Run starts the application with the given configuration
func Run(app config.Application) (err error) {
	err = app.Validate()
	if err != nil {
		return
	}
	// for debugging

	dbConn := db.New(app.Dialect, app.Dsn, app.Debug)
	if app.Migrate {
		err = db.Migrate(dbConn, &entity.Port{})
		if err != nil {
			return
		}
	}

	portUseCase := usecase.NewPortUseCase(repo.NewPortRepo(dbConn))

	http.InitUploadLimits(app.MaxFileSize, app.MemoryLimit)

	controller := http.NewController(portUseCase, app.WorkerCount, app.FileField, app.Path)

	httpServer := httpserver.New(http.Recovery(controller), httpserver.Port(fmt.Sprintf("%d", app.Port)))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	return
}
