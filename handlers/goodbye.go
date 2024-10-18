package handlers

import (
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye!")
	_, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error!", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Goodbye!"))
}
