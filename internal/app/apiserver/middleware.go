package apiserver

import (
	"context"
	"net/http"
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
			//s.error(w, http.StatusUnauthorized, errNotAuthenticated)
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			//s.error(w, http.StatusUnauthorized, errNotAuthenticated)
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) isLoggedOFF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		_, ok := session.Values["user_id"]

		if ok {
			//s.respond(w, http.StatusOK, "already logged in")
			http.Redirect(w, r, "/whoami", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
