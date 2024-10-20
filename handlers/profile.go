package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid url", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "Invalid url", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			http.Error(w, "Cannot decode id", http.StatusInternalServerError)
			return
		}

		p.updateProfiles(id, w, r)
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

func (p *Profile) updateProfiles(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")

	prof := &data.User{}
	err := prof.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Cannot unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prof: %#v", prof)
	err = data.UpdateUser(id, prof)

	if err == data.ErrorUserNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

}
