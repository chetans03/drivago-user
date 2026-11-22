package runner

import (
	"fmt"
	"sync"

	"github.com/chetan03/drivago-user/internal/config"
	"github.com/chetan03/drivago-user/internal/route"
)

func RunApplication() error {

	wg := &sync.WaitGroup{}

	if err := config.InitializeLogging(); err != nil {
		return fmt.Errorf("initialize logging: %w", err)
	}

	config.Logger.Info("F02: Starting the application")

	if err := config.LoadEnvVars(); err != nil {
		config.Logger.WithError(err).Error("load configuration")
		return fmt.Errorf("load configuration: %w", err)
	}

	if err := config.InitializeDatabase(); err != nil {
		config.Logger.WithError(err).Error("initialize database")
		return fmt.Errorf("initialize database: %w", err)
	}

	router := route.NewRoute()
	route.RegisterRoute(router)
	srv := config.NewServer(router)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.StartServer(); err != nil {
			config.Logger.Fatalf("Error starting server: %w", err)
		}
	}()

	wg.Wait()

	config.Logger.Info("Application startup completed")
	return nil
}
