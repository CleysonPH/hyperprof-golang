package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	r.Body = b
	return r.ResponseWriter.Write(b)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		handler.ServeHTTP(rec, r)
		duration := time.Since(startTime)

		logger := log.Info()
		if rec.StatusCode >= 400 {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.Str("protocol", r.Proto).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Dur("duration", duration).
			Msg("received a request")
	})
}
