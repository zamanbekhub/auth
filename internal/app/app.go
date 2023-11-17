package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service/internal/config"
	httpDelivery "service/internal/delivery/http"
	"service/internal/repository"
	httpServer "service/internal/server"
	"service/internal/service"
	"service/pkg/db/postgresql"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	postgresDB, err := postgresql.NewDB(cfg.Database.PostgreDSN)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepositories(postgresDB)

	services := service.NewServices(repos)

	handler := httpDelivery.NewHandlerDelivery(logger, services, "service")

	srv, err := httpServer.NewServer(cfg, handler)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("🔥 Server stopped due error", "err", err.Error())
		} else {
			logger.Printf("✅ Server shutdown successfully")
		}
	}()

	logger.Printf("🚀 Starting server at http://0.0.0.0:%s", cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer shutdownCtxCancel()

	isShutdownErrors := false

	if err = srv.Shutdown(shutdownCtx); err != nil {
		logger.Printf(err.Error())
		isShutdownErrors = true
	}

	if isShutdownErrors {
		logger.Printf("Server closed, but not all resources closed properly!")
	} else {
		logger.Printf("✅ Server shutdown successfully")
	}
}
