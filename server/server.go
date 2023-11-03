package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/spy16/wisenotes/storage/db"
)

var (
	errBadRequest = errors.New("bad request")
	errNotFound   = errors.New("not found")
)

// Serve starts the server at given address.
func Serve(ctx context.Context, addr string, queries *db.Queries, ui http.Handler) error {
	r := chi.NewRouter()

	r.Mount("/", ui)

	r.Route("/api", func(r chi.Router) {
		r.Get("/profiles", getProfiles(queries))
		r.Post("/profiles", createProfile(queries))

		r.Get("/profiles/{profileID}/articles", listArticles(queries))
		r.Post("/profiles/{profileID}/articles", createArticle(queries))
	})

	return serveCtx(ctx, addr, r)
}

// ServeCtx starts an HTTP server and blocks until the context is canceled.
// Context cancellation triggers a graceful shutdown of the server.
func serveCtx(ctx context.Context, addr string, handler http.Handler) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return err

	case <-ctx.Done():
		log.Debug().Msg("shutting down server")
		return srv.Shutdown(ctx)
	}
}

// sendJSON sends the given value as JSON response.
func sendJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Warn().Err(err).Msg("failed to send response")
	}
}

type httpHandlerE func(w http.ResponseWriter, r *http.Request) error

func httpHandler(h httpHandlerE) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			switch {
			case errors.Is(err, errBadRequest):

				sendJSON(w, http.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})

			case errors.Is(err, sql.ErrNoRows), errors.Is(err, errNotFound):
				sendJSON(w, http.StatusNotFound, map[string]string{
					"error": err.Error(),
				})

			default:
				log.Error().Err(err).Msg("unexpected error")
				sendJSON(w, http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
		}
	}
}
