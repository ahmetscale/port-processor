package http

import (
	"net/http"
)

// UploadHandler is the handler for the upload endpoint
func (c Controller) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxFileSize))
	if err := r.ParseMultipartForm(int64(memoryLimit / 2)); err != nil {
		http.Error(w, "memory error", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile(c.FileField)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	err = c.p.Process(r.Context(), file, c.WorkerCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
