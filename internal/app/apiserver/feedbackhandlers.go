package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/gorilla/mux"
)

func (s *server) handleFeedbackBook() http.HandlerFunc {
	type request struct {
		Feedback string `json:"feedback"`
		Grade    int    `json:"grade"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		book_id, err := strconv.Atoi(mux.Vars(r)["book_id"])
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		ID := r.Context().Value(ctxKeyUser).(*model.User).ID
		f := &model.Feedback{
			UserID:   ID,
			BookID:   book_id,
			Feedback: req.Feedback,
			Grade:    req.Grade,
		}

		if err := s.store.Feedback().Create(f); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		/* maybe need transactions create feedback + update grade */

		/*if err := s.store.Book().UpdateGrade(book_id); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		} */

		s.respond(w, http.StatusOK, f)
	}
}

func (s *server) handleGetFeedbackBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book_id, err := strconv.Atoi(mux.Vars(r)["book_id"])
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		feedback, err := s.store.Feedback().FindByBook(book_id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, feedback)
	}
}
