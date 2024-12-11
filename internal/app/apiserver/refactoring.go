package apiserver

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (s *server) handleLoadBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(10 << 20)

		bookID := r.FormValue("bookId")
		if bookID == "" {
			s.error(w, http.StatusBadRequest, errors.New("book ID is required"))
			return
		}
		ID, err := strconv.Atoi(bookID)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		file, handler, err := r.FormFile("book")
		if err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}
		defer file.Close()

		contentType := handler.Header.Get("Content-Type")

		switch contentType {
		case "application/pdf":
			contentType = "pdf"
		case "application/epub+zip":
			contentType = "epub"
		case "application/x-mobipocket-ebook":
			contentType = "mobi"
		default:
			s.error(w, http.StatusBadRequest, fmt.Errorf("unsupported file type: %s", contentType))
			return
		}

		if _, err := s.store.Book().FindByID(ID); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Book().UpdateFile(ID, contentType); err != nil {
			s.logger.Info(err)
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.logger.Info("MIME Header: ", handler.Header)

		dst, err := os.Create(filepath.Join(basePath, fmt.Sprintf("./loads/%s.%s", strconv.Itoa(ID+1000), contentType)))
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, "Successfully Uploaded File")
	}
}

func (s *server) handleUploadBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := mux.Vars(r)
		book_id := mp["book_id"]
		id, err := strconv.Atoi(book_id)
		id += 1000
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		dir := filepath.Join(basePath, "/loads")
		files, err := os.ReadDir(dir)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		var foundFile string
		for _, file := range files {
			if strings.HasPrefix(file.Name(), strconv.Itoa(id)+".") {
				foundFile = filepath.Join(dir, file.Name())
				break
			}
		}

		if foundFile == "" {
			s.error(w, http.StatusNotFound, errors.New("file not found"))
			return
		}

		/*file, err := os.Open(foundFile)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		defer file.Close()

		ext := filepath.Ext(foundFile)

		switch ext {
		case ".pdf":
			w.Header().Set("Content-Type", "application/pdf")
		case ".epub":
			w.Header().Set("Content-Type", "application/epub+zip")
		case ".mobi":
			w.Header().Set("Content-Type", "application/x-mobipocket-ebook")
		case ".txt":
			w.Header().Set("Content-Type", "text/plain")
		case ".html", ".htm":
			w.Header().Set("Content-Type", "text/html")
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		default:
			w.Header().Set("Content-Type", "application/octet-stream")
		}

		w.Header().Set("Content-Disposition", "inline; filename=\""+filepath.Base(foundFile)+"\"")

		if _, err := io.Copy(w, file); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}*/

		ext := filepath.Ext(foundFile)
		if ext != ".pdf" && ext != ".epub" {
			s.error(w, http.StatusUnsupportedMediaType, errors.New("unsupported file type"))
			return
		}

		data := map[string]interface{}{
			"FileName": filepath.Base(foundFile),
			"FileType": ext,
			"FilePath": foundFile,
		}

		s.executeTemplate(w, data, filepath.Join(basePath, "/static/read.html"))
	}
}
