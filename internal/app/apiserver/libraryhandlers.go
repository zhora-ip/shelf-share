package apiserver

import (
	"net/http"
	"path/filepath"

	"github.com/zhora-ip/shelf-share/internal/app/model"
	"github.com/zhora-ip/shelf-share/internal/app/store"
)

func (s *server) handleGetLibrary() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(ctxKeyUser).(*model.User).ID

		library, err := s.store.Library().FindAll(id)
		if err != nil {
			if err == store.ErrRecordNotFound {
				s.respond(w, http.StatusNotFound, "library is empty")
				return
			}

			s.error(w, http.StatusInternalServerError, err)
			return
		}

		books := make([]*model.Book, 0)

		for _, v := range library {
			book, err := s.store.Book().FindByID(v)
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			books = append(books, book)
		}

		//s.respond(w, http.StatusOK, books)
		s.executeTemplate(w, books, filepath.Join(basePath, "/static/library.html"))
	}
}
