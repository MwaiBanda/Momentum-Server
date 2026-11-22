package main

import (
	"Momentum/application"
	_ "Momentum/docs"

	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title			Momentum API
// @version		v1
// @description	This is a backend service to process payments and provide some CRUD functionality
// @termsOfService	https://momentumindiana.org/statement-of-faith/
// @contact.name	Mwai Banda
// @contact.email	bandamwai@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			services.momentumchurch.dev
// @BasePath		/
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Bearer token required for protected API routes.

func main() {
	if err := run(); err != nil {
		log.Printf("Momentum API stopped: %v", err)
		os.Exit(1)
	}
}

func run() (runErr error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	startupContext, cancelStartup := context.WithTimeout(ctx, 15*time.Second)
	defer cancelStartup()

	log.Println("Starting Momentum API")
	app, err := application.New(startupContext)
	if err != nil {
		return fmt.Errorf("initialize application: %w", err)
	}
	defer func() {
		if err := app.Close(); err != nil {
			runErr = errors.Join(runErr, err)
		}
	}()

	if err := app.Run(ctx); err != nil {
		return fmt.Errorf("run application: %w", err)
	}
	return nil
}
