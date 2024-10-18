package handlers

import (
	"log"
	"net/http"

	"github.com/ZhoraIp/ShelfShare/data"
)

type Profile struct {
	l *log.Logger
}

func NewProfile(l *log.Logger) *Profile {
	return &Profile{l}
}

func (p *Profile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProfiles(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProfiles(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Profile) getProfiles(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Profiles")
	pf := data.GetProfiles()
	err := pf.ToJSON(w)

	if err != nil {
		http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
	}
}

func (p *Profile) addProfiles(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Profiles")

	prof := &data.User{}
	err := prof.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Cannon unmarshal json", http.StatusInternalServerError)
	}

	p.l.Printf("Prof: %#v", prof)

	data.AddUser(prof)
}
