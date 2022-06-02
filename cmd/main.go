package main

import (
	"port-processor/internal/app"
	"port-processor/pkg/config"
)

func main() {
	cfg, err := config.GetApplication()
	if err != nil {
		panic(err)
	}
	panic(app.Run(cfg))
}
