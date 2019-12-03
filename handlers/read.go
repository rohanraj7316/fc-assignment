package handlers

import (
	"bufio"
	"fmt"
	"freecharge/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// ReadFile read the data from a file.
func ReadFile(w http.ResponseWriter, r *http.Request) error {

	fnVal := chi.URLParam(r, "fileName")

	if AccessBody[fnVal] {

		fd, err := utils.GetAudioFileReadDiscriptor(fnVal)
		if err != nil {
			return err
		}
		defer func() {
			if err := fd.Close(); err != nil {
				log.Println(err)
			}
		}()

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Accept-Ranges", "bytes")

		re := bufio.NewReader(fd)
		b := make([]byte, 30)
		for {
			n, err := re.Read(b)
			if err != nil {
				log.Println(err)
				break
			}
			w.Write(b[0:n])
		}
	} else {
		return fmt.Errorf("error: access disabled")
	}
	return nil
}
