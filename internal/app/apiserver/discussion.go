package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleCreateDiscussion() http.HandlerFunc {

	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
		}

		d := &model.Discussion{
			UserID:      r.Context().Value(ctxKeyUser).(*model.User).ID,
			Title:       req.Title,
			Description: req.Description,
		}

		if err := s.store.Discussion().Create(d); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, d)
	}
}

func (s *server) handleMessageDiscussion() http.HandlerFunc {

	type request struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		discussionID, ok := mux.Vars(r)["discussion_id"]
		if !ok {
			s.error(w, http.StatusInternalServerError, errors.New("cannot find discussion id"))
			return
		}

		ID, err := strconv.Atoi(discussionID)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		m := &model.Message{
			UserID:       r.Context().Value(ctxKeyUser).(*model.User).ID,
			DiscussionID: ID,
			Message:      req.Message,
		}

		if err := s.store.Discussion().NewMessage(m); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, m)
	}
}

func (s *server) handleGetDiscussion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := mux.Vars(r)["discussion_id"]
		if !ok {
			s.error(w, http.StatusInternalServerError, errors.New("discussion_id not found"))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		d, ms, err := s.store.Discussion().FindByID(idInt)

		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		res := struct {
			Discussion *model.Discussion
			Messages   []*model.Message
		}{
			Discussion: d,
			Messages:   ms,
		}

		//s.respond(w, http.StatusOK, res)
		s.executeTemplate(w, res, filepath.Join(basePath, "/static/forum.html"))
	}
}
