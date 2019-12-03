package handlers

import (
	"freecharge/utils"
	"io"
	"net/http"
)

// UploadFile used for uploading file.
func UploadFile(w http.ResponseWriter, r *http.Request) error {

	fd, err := utils.GetAudioFileWriteDiscriptor()
	if err != nil {
		return err
	}
	defer func() {
		fd.Close()
	}()

	// writing content into io copy.
	_, err = io.Copy(fd, r.Body)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	return nil
}
