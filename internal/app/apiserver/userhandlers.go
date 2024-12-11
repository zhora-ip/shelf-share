package apiserver

import (
	"net/http"
	"path/filepath"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
)

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ctxKeyUser).(*model.User)

		if !ok || user == nil {
			s.error(w, http.StatusUnauthorized, nil)
			return
		}

		//s.respond(w, http.StatusOK, user)

		s.executeTemplate(w, user, filepath.Join(basePath, "/static/whoami.html"))
	}
}

func (s *server) handleGetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.User().FindAll()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		//s.respond(w, http.StatusOK, us)
		s.executeTemplate(w, users, filepath.Join(basePath, "/static/users.html"))
	}
}
