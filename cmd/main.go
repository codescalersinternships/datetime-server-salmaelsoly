package main

import (
	"context"
	"errors"
	"httpServer/api"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.Handle("/datetime", http.HandlerFunc(routes.DateTime))
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Any("HTTP server error: %v", err)
		}
		slog.Warn("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Any("HTTP shutdown error: %v", err)
	}
	slog.Info("Graceful shutdown complete.")

}
