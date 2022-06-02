package http

import (
	"fmt"
	"net/http"
	"port-processor/internal/usecase"
)

type Controller struct {
	p           usecase.Port
	WorkerCount int
	FileField   string
	Path        string
}

func NewController(p usecase.Port, workerCount int, fileField string, uploadPath string) *Controller {
	return &Controller{p: p, WorkerCount: workerCount, FileField: fileField, Path: uploadPath}
}

const MegaByte = 1024 * 1024

var maxFileSize = MegaByte * 1024 // default max upload size is 1GB

var memoryLimit = MegaByte * 100 // default memory limit is 100MB

// InitUploadLimits sets the max upload size and memory limit
func InitUploadLimits(memoryLimit, fileSizeLimit int) {
	memoryLimit = MegaByte * memoryLimit
	maxFileSize = MegaByte * fileSizeLimit
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if req.URL.Path == c.Path {
			c.UploadHandler(w, req)
		} else {
			http.Error(w, fmt.Sprintf("please use %s path", c.Path), http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
