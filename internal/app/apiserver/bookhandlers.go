package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/ZhoraIp/ShelfShare/internal/app/model"
	"github.com/gorilla/mux"
)

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

		ID := r.Context().Value(ctxKeyUser).(*model.User).ID
		b := &model.Book{
			Title:       req.Title,
			Author:      req.Author,
			Genre:       req.Genre,
			Description: req.Description,
			CreatedBy:   ID,
		}

		if err := s.store.Book().Create(b); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, b)
	}
}

func (s *server) handleAddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId := r.Context().Value(ctxKeyUser).(*model.User).ID

		vars := mux.Vars(r)
		id := vars["book_id"]
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

func (s *server) handleGetBooks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		author := query.Get("author")
		title := query.Get("title")
		genre := query.Get("genre")

		books, err := s.store.Book().FindAll()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		result := make([]*model.Book, 0)

		for _, v := range books {
			if (author == "" || author == v.Author) && (title == "" || title == v.Title) && (genre == "" || genre == v.Genre) {
				result = append(result, v)
			}
		}

		//s.respond(w, http.StatusOK, bs)
		s.executeTemplate(w, result, filepath.Join(basePath, "/static/library.html"))
	}
}

func (s *server) handleGetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := mux.Vars(r)
		book_id, ok := mp["book_id"]
		if !ok {
			s.error(w, http.StatusInternalServerError, errors.New("cannot fid book_id"))
			return
		}
		id, err := strconv.Atoi(book_id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		b, err := s.store.Book().FindByID(id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		//s.respond(w, http.StatusOK, b)
		s.executeTemplate(w, b, filepath.Join(basePath, "static/book.html"))
	}
}
