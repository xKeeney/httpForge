package httpMiddlewares

import (
	"bytes"
	"io"
	"net/http"

	"github.com/xKeeney/httpForge/httpLogger"

	"time"
)

// statusRecorder запоминает код ответа
type statusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader сохраняет код и вызывает оригинальный метод
func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Write переопределён, чтобы корректно обрабатывать случай,
// когда WriteHeader не был вызван явно (статус по умолчанию 200).
func (r *statusRecorder) Write(b []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.ResponseWriter.Write(b)
}

type baseMiddlewares struct {
	logger *httpLogger.HttpLogger
}

func InitBaseMiddlewares(logger *httpLogger.HttpLogger) *baseMiddlewares {
	return &baseMiddlewares{logger: logger}
}

func (m *baseMiddlewares) InfoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK, // начальное значение, но может быть перезаписано
		}
		start := time.Now()
		m.logger.Printf("[INFO:RECIVE]    %s - \"%s %s\"\n",
			r.RemoteAddr, r.Method, r.URL.Path,
		)

		next.ServeHTTP(recorder, r)

		m.logger.Printf("[INFO:COMPLETE]  %s - \"%s %s\" -> status %d in %v\n",
			r.RemoteAddr, r.Method, r.URL.Path, recorder.status, time.Since(start),
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
