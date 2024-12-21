package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ZhoraIp/ShelfShare/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "session"
	ctxKeyUser  ctxKey = iota
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	//errNotAuthenticated         = errors.New("not authenticated")
	basePath = "/"
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func NewServer(store store.Store, sessionStore sessions.Store) *server {

	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Shutdown(ctx context.Context) error {
	return nil
}

func (s *server) configureRouter() {

	execPath, err := os.Executable()
	if err != nil {
		s.logger.Error("Error getting executable path:", err)
		return
	}
	basePath = filepath.Dir(execPath)

	s.router.HandleFunc("/", s.handleStart())

	s.configurePublicRoutes()
	s.configurePrivateRoutes()
}

func (s *server) configurePublicRoutes() {
	r := s.router.NewRoute().Subrouter()
	r.Use(s.isLoggedOFF)

	r.HandleFunc("/registration", s.handleGetRegistration()).Methods(http.MethodGet)
	r.HandleFunc("/registration", s.handleRegistration()).Methods(http.MethodPost)
	r.HandleFunc("/login", s.handleGetLogin()).Methods(http.MethodGet)
	r.HandleFunc("/login", s.handleLogin()).Methods(http.MethodPost)
}

func (s *server) configurePrivateRoutes() {
	r := s.router.NewRoute().Subrouter()
	r.Use(s.authUser)

	s.configureBooksRoutes(r)
	s.configeureDiscussionRoutes(r)

	r.HandleFunc("/logoff", s.handleLogoff()).Methods(http.MethodPost)
	r.HandleFunc("/logoff", s.handleGetLogoff()).Methods(http.MethodGet)

	r.HandleFunc("/whoami", s.handleWhoAmI()).Methods(http.MethodGet)
	r.HandleFunc("/users", s.handleGetUsers()).Methods(http.MethodGet)

	r.HandleFunc("/library/books/{book_id:[0-9]+}", s.handleAddBook()).Methods(http.MethodPost)
	r.HandleFunc("/library", s.handleGetLibrary()).Methods(http.MethodGet)

}

func (s *server) configureBooksRoutes(r *mux.Router) {
	r.HandleFunc("/books", s.handleGetBooks()).Methods(http.MethodGet)
	r.HandleFunc("/books/{book_id:[0-9+]}", s.handleGetBook()).Methods(http.MethodGet)
	r.HandleFunc("/books/creation", s.handleCreateBooks()).Methods(http.MethodPost)
	r.HandleFunc("/books/loading", s.handleGetLoadBook()).Methods(http.MethodGet)
	r.HandleFunc("/books/loading", s.handleLoadBook()).Methods(http.MethodPatch)
	r.HandleFunc("/books/{book_id:[0-9]+}/feedback", s.handleFeedbackBook()).Methods(http.MethodPost)
	r.HandleFunc("/books/{book_id:[0-9]+}/feedback", s.handleGetFeedbackBook()).Methods(http.MethodGet)
	r.HandleFunc("/books/{book_id:[0-9]+}/uploading", s.handleUploadBook()).Methods(http.MethodGet)

	r.PathPrefix("/loads/").Handler(http.StripPrefix("/loads/", http.FileServer(http.Dir("loads"))))

}

func (s *server) configeureDiscussionRoutes(r *mux.Router) {
	r.HandleFunc("/discussion", s.handleGetCreateDiscussion()).Methods(http.MethodGet)
	r.HandleFunc("/discussion", s.handleCreateDiscussion()).Methods(http.MethodPost)
	r.HandleFunc("/discussion/{discussion_id:[0-9]+}", s.handleMessageDiscussion()).Methods(http.MethodPost)
	r.HandleFunc("/discussion/{discussion_id:[0-9]+}", s.handleGetDiscussion()).Methods(http.MethodGet)

}

func (s *server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
