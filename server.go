package httpForge

import (
	"fmt"
	"net/http"

	"github.com/xKeeney/httpForge/httpLogger"
)

type HttpApp struct {
	server      *http.Server
	rootMux     *http.ServeMux
	middlewares []func(http.Handler) http.Handler
	logger      *httpLogger.HttpLogger
}

func NewHttpApp(addr string, logger *httpLogger.HttpLogger) *HttpApp {
	rootMux := http.NewServeMux()
	s := &http.Server{
		Addr:    addr,
		Handler: rootMux,
	}
	return &HttpApp{
		server:      s,
		rootMux:     rootMux,
		middlewares: make([]func(http.Handler) http.Handler, 0),
		logger:      logger,
	}
}

// AddMiddleware добавляет middleware для всего приложения
func (a *HttpApp) AddMiddleware(middleware func(http.Handler) http.Handler) {
	a.middlewares = append(a.middlewares, middleware)
}

func (a *HttpApp) Get(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	pattern := fmt.Sprintf("GET %s", route)
	a.rootMux.HandleFunc(pattern, handler)
}

func (a *HttpApp) Post(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	pattern := fmt.Sprintf("POST %s", route)
	a.rootMux.HandleFunc(pattern, handler)
}

func (a *HttpApp) Put(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	pattern := fmt.Sprintf("PUT %s", route)
	a.rootMux.HandleFunc(pattern, handler)
}

func (a *HttpApp) Delete(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	pattern := fmt.Sprintf("DELETE %s", route)
	a.rootMux.HandleFunc(pattern, handler)
}

func (a *HttpApp) ListenAndServe() {
	a.logger.Printf("[INFO]	Server starting...\n")
	var handler http.Handler = a.rootMux
	// Применяем middleware в обратном порядке
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		handler = a.middlewares[i](handler)
	}
	a.server.Handler = handler
	a.logger.Printf("[INFO]	Server started at \"%s\"\n", a.server.Addr)
	a.server.ListenAndServe()
}
