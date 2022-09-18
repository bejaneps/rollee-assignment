package transport

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Service represents the service which provides
// Most Frequent Word functionality
type Service interface {
	// UpsertWord updates frequency of word count
	// or inserts word in the storage
	UpsertWord(word string) error

	// GetMostFrequentWord returns most frequest
	// word stored in the service
	GetMostFrequentWord(beginning string) (string, error)
}

// Handler is http Most Frequent Word handler
type Handler struct {
	service Service
	server  *http.Server
}

// New instantiates an instance of MFW http Handler
func New(service Service, port string, opts ...Option) *Handler {
	handler := &Handler{
		service: service,
	}

	handler.initServer(port)

	// apply options
	for _, opt := range opts {
		opt(handler)
	}

	return handler
}

func (h *Handler) initServer(port string) {
	router := mux.NewRouter()

	// POST /service/word="something"
	router.Handle(
		`/service/word="{word:[a-zA-Z]+}"`,
		handleServicePost(h.service),
	).Methods(http.MethodPost)

	// GET /service/word="something"
	router.Handle(
		`/service/word="{word:[a-zA-Z]+}"`,
		handleServiceGet(h.service),
	).Methods(http.MethodGet)

	// Not found handler
	router.NotFoundHandler = handleNotFound()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	h.server = server
}

// Serve listens and serves http
func (h *Handler) Serve() error {
	log.Printf("starting server on %s\n", h.server.Addr)

	return h.server.ListenAndServe()
}

// Close shutdowns gracefully http server
func (h *Handler) Close(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

// Option is option for MFW Handler.
type Option func(*Handler)

// WithReadTimeout sets request read timeout in seconds.
func WithReadTimeout(timeout int64) Option {
	return func(h *Handler) {
		h.server.ReadTimeout = time.Duration(timeout) * time.Second
	}
}

// WithWriteTimeout sets response write timeout in seconds.
func WithWriteTimeout(timeout int64) Option {
	return func(h *Handler) {
		h.server.WriteTimeout = time.Duration(timeout) * time.Second
	}
}
