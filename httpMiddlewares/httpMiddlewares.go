package httpMiddlewares

import (
	"bytes"
	"io"
	"net/http"

	"github.com/xKeeney/httpForge/httpLogger"

	"time"
)

type baseMiddlewares struct {
	logger *httpLogger.HttpLogger
}

func InitBaseMiddlewares(logger *httpLogger.HttpLogger) *baseMiddlewares {
	return &baseMiddlewares{logger: logger}
}

func (m *baseMiddlewares) InfoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.logger.Printf("[INFO:RECIVE]    %s - \"%s %s\"\n",
			r.RemoteAddr, r.Method, r.URL.Path,
		)

		next.ServeHTTP(w, r)

		m.logger.Printf("[INFO:COMPLETE]  %s - \"%s %s\" in %v\n",
			r.RemoteAddr, r.Method, r.URL.Path, time.Since(start),
		)
	})
}

func (m *baseMiddlewares) RequestsLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			m.logger.Tracef("Error while read body: %v", err)
			http.Error(w, "Bad request body", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		m.logger.Tracef(
			"Request details:\nMethod: %s\nPath: %s\nBody: %s\n",
			r.Method,
			r.URL.Path,
			string(bodyBytes),
		)

		next.ServeHTTP(w, r)
	})
}
