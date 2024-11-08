package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

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
	errNotAuthenticated         = errors.New("not authenticated")
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

func (s *server) configureRouter() {

	public := s.router.NewRoute().Subrouter()
	public.Use(s.isLoggedOff)

	public.HandleFunc("/registration", s.handleRegistration()).Methods(http.MethodPost)
	public.HandleFunc("/login", s.handleLogin()).Methods(http.MethodPost)

	protected := s.router.PathPrefix("/session").Subrouter()
	protected.Use(s.authUser)
	protected.HandleFunc("/logoff", s.handleLogoff()).Methods(http.MethodPost)

	protected.HandleFunc("/whoami", s.handleWhoAmI()).Methods(http.MethodGet)
	protected.HandleFunc("/users", s.handleGetUsers()).Methods(http.MethodGet)

	protected.HandleFunc("/books", s.handleGetBooks()).Methods(http.MethodGet)
	protected.HandleFunc("/books/creation", s.handleCreateBooks()).Methods(http.MethodPost)

	private := s.router.PathPrefix("/session").Subrouter()
	private.Use(s.authUser, s.checkUser)
	private.HandleFunc("/users/{id:[0-9]+}/library/books/{book_id:[0-9]+}", s.handleCreateBook()).Methods(http.MethodPost)
	private.HandleFunc("/users/{id:[0-9]+}/library", s.handleGetLibrary()).Methods(http.MethodGet)
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
