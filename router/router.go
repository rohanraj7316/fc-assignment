package router

import (
	"freecharge/handlers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// format for handler.
type handler func(w http.ResponseWriter, r *http.Request) error

// New create the new server.
func New(port int) error {

	mux := handleRouter()

	server := http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("listening on port: " + strconv.Itoa(port))
	panic(server.ListenAndServe())
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Auth-Token")
		if auth == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			if auth == "1234567" {
				next.ServeHTTP(w, r)
			} else {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
	})
}

func handleRouter() *chi.Mux {

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS", "POST"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.NoCache)

	r.Get("/health", h(handlers.GetHealth))

	r.Route("/get", func(r chi.Router) {
		r.Get("/{fileName}", h(handlers.ReadFile))
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(auth)
		r.Post("/upload_file", h(handlers.UploadFile))
		r.Post("/update_file_access", h(handlers.UpdateAccess))
		r.Delete("/delete_file/{fileName}", h(handlers.DeleteFile))
	})

	return r
}

// ServerHTTP check for the errors from all handlers.
func (fn handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func h(fn handler) http.HandlerFunc {
	return handler(fn).ServerHTTP
}
