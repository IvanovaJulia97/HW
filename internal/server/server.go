package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// тут создаем роутер
func MyServer(logger *log.Logger) *Server {
	rout := chi.NewRouter()

	rout.Get("/", handlers.ReadHandler)
	rout.Post("/upload", handlers.EditHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}

}

func (s *Server) Start() error {
	s.Logger.Println("Сервер запущен")

	err := s.HTTPServer.ListenAndServe()
	if err != nil {
		s.Logger.Println("Ошибка запуска сервера", err)
	}
	return err

}
