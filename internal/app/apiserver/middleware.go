package apiserver

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))

	})
}

func (s *server) isLoggedOff(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		_, ok := session.Values["user_id"]

		if ok {
			s.respond(w, http.StatusOK, "already logged in")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *server) checkUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u := r.Context().Value(ctxKeyUser).(*model.User)
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if u.Id != id {
			s.respond(w, http.StatusForbidden, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
