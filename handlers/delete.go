package handlers

import (
	"fmt"
	"freecharge/utils"
	"net/http"

	"github.com/go-chi/chi"
)

// DeleteFile used to delete file.
func DeleteFile(w http.ResponseWriter, r *http.Request) error {

	fnVal := chi.URLParam(r, "fileName")

	if fnVal == "" {
		return fmt.Errorf("error in getting filename")
	}

	err := utils.DelFile(fnVal)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	return nil
}
