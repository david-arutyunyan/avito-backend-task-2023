package avito

import (
	"context"
	"net/http"
	"time"
)

type Server struct { // Структура сервера для запуска HTTP сервера
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error { // Запуск сервера
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe() // Бесконечный цикл for, который слушает все запросы для последующей обработка
}

func (s *Server) Shutdown(ctx context.Context) error { // Остановка сервера
	return s.httpServer.Shutdown(ctx)
}
