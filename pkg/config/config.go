package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

// Application holds application configurations
type Application struct {
	Dsn         string `required:"true" env:"DSN"`
	Port        int    `default:"5050" env:"PORT"`
	MemoryLimit int    `default:"200" env:"MEMORY_LIMIT"`   // in MB, default is 200MB
	MaxFileSize int    `default:"1024" env:"MAX_FILE_SIZE"` // in MB, default is 1GB
	Debug       bool   `env:"PORT_DEBUG_MODE"`
	Migrate     bool   `env:"PORT_MIGRATE" default:"true"`
	WorkerCount int    `default:"1" env:"WORKER_COUNT"`
	FileField   string `env:"FILE_FIELD" default:"file"`
	Path        string `env:"UPLOAD_PATH" default:"/upload"` // upload Path
	Dialect     string `env:"DIALECT" default:"postgres"`
}

func (a Application) Validate() error {
	if a.Port <= 0 {
		return fmt.Errorf("invalid port")
	}
	if a.MemoryLimit <= 0 {
		return fmt.Errorf("invalid memory limit")
	}
	if a.MaxFileSize <= a.MemoryLimit {
		return fmt.Errorf("max file size must be greater than memory limit")
	}
	if a.WorkerCount <= 0 {
		return fmt.Errorf("invalid worker count")
	}
	if a.WorkerCount > 16 {
		return fmt.Errorf("worker count should be less than 16")
	}
	return nil
}

func GetApplication() (app Application, err error) {
	err = configor.Load(&app)
	return
}
