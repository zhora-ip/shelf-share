package apiserver

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/start.html"))
	}
}

func (s *server) handleGetRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/registration.html"))
	}
}

func (s *server) handleGetLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/login.html"))
	}
}

func (s *server) handleGetLogoff() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/logoff.html"))
	}
}

func (s *server) handleGetLoadBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/load.html"))
	}
}
func (s *server) handleGetCreateDiscussion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.executeTemplate(w, nil, filepath.Join(basePath, "/static/discussion.html"))
	}
}

func (s *server) executeTemplate(w http.ResponseWriter, data interface{}, filepath string) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		s.error(w, http.StatusInternalServerError, err)
	}

}
