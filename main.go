package main

import (
	"sync"

	"github.com/urcop/emotionalTracker/app"
	"github.com/urcop/emotionalTracker/services/logger"
)

var wg sync.WaitGroup

func main() {
	logger := logger.Init("dev")
	logger.Info("Starting application")
	wg.Add(2)
	go func() {
		defer wg.Done()
		logger.Info("Starting HTTP server")
		app.NewHttpServer().Start()
		logger.Info("HTTP server stopped")
	}()

	go func() {
		defer wg.Done()
		logger.Info("Starting GRPC server")
		app.NewGRPCServer().Start()
		logger.Info("GRPC server stopped")
	}()
	wg.Wait()
	logger.Info("Application shutdown")
}
