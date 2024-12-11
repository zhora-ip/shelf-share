package apiserver

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZhoraIp/ShelfShare/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

	srv.logger.Info("starting api server")

	shutdownCtx, stopServer := context.WithCancel(context.Background())
	defer stopServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	listenAndServeErrChan := make(chan error, 1)

	go func() {
		listenAndServeErrChan <- http.ListenAndServe(config.BindAddr, srv)
	}()

	select {
	case <-sigChan:
		srv.logger.Info("Recieved terminate, graceful shutdown!")
	case err := <-listenAndServeErrChan:
		if err != nil && err != http.ErrServerClosed {
			srv.logger.Fatalf("Recieved error from listen and serve, %v", err)
			return err
		}
	}

	ctx, cancel := context.WithTimeout(shutdownCtx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		srv.logger.Fatalf("Server Shutdown Error: %v", err)
		return err
	}

	srv.logger.Info("Server exited gracefully")
	return nil
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
