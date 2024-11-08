package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
)

func (s *server) handleRegistration() http.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Nickname: req.Nickname,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, http.StatusCreated, u)
	}
}

func (s *server) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.Id

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, "succesfully logged in")
	}
}

func (s *server) handleLogoff() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		session.Options.MaxAge = -1
		delete(session.Values, "user_id")

		if err := session.Save(r, w); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusNoContent, "succesfully logged off")
	}
}
