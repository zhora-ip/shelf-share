package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/ZhoraIp/ShelfShare/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	/* graceful shutdown

	go func() {
		err := s.server.ListenAndServe()
		if err == http.ErrServerClosed {
			s.logger.Info(err)
		} else if err != nil {
			s.logger.Error(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	fmt.Println("Recieved terminate, graceful shutdown!", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.server.Shutdown(tc)
	*/

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

	srv.logger.Info("starting api server")

	return http.ListenAndServe(config.BindAddr, srv)
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
