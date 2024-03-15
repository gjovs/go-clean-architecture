package www

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: port,
	}
}

func (server *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	server.Handlers[path] = handler
}

func (server *WebServer) Start() {
	server.Router.Use(middleware.Logger)
	
	for path, handler := range server.Handlers {
		server.Router.Handle(path, handler)
	}

	http.ListenAndServe(server.WebServerPort, server.Router)
}
