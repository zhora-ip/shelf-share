package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/ZhoraIp/ShelfShare/internal/app/store"
	"github.com/gorilla/mux"
)

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleGetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := s.store.User().FindAll()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, http.StatusOK, us)

	}
}

func (s *server) handleCreateBooks() http.HandlerFunc {

	type request struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Genre       string `json:"genre"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		b := &model.Book{
			Title:       req.Title,
			Author:      req.Author,
			Genre:       req.Genre,
			Description: req.Description,
		}

		if err := s.store.Book().Create(b); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, b)
	}
}

func (s *server) handleGetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := s.store.Book().FindAll()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, http.StatusOK, bs)
	}
}

func (s *server) handleCreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		userId, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		id = vars["book_id"]
		bookId, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Library().AddBook(userId, bookId); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
	}
}

func (s *server) handleGetLibrary() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		userId, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		library, err := s.store.Library().FindAll(userId)
		if err != nil {
			if err == store.ErrRecordNotFound {
				s.respond(w, http.StatusNotFound, "library is empty")
				return
			}

			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, library)
	}
}
