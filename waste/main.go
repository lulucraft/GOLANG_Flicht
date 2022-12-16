package main

import (
	"log"
	"net/http"

	"git.evoliatis.fr/scollado/system"
	"github.com/gorilla/mux"
)

func HTTPippasserelle(w http.ResponseWriter, _ *http.Request) {
	log.Print("/ip/passerelle")
	m, err := system.GetPasserelle()
	if err != nil {
		log.Printf("Feeler /ip/passerelle : %v", err)
	}
	system.EnvoiJSON(m, "/ip/passerelle", w)
}

func HTTPip(w http.ResponseWriter, _ *http.Request) {
	log.Print("/ip")
	m, err := system.GetAdresseIP()
	if err != nil {
		log.Printf("Feeler /ip : %v", err)
	}
	system.EnvoiJSON(m, "/ip", w)
}

func HTTP1ip(w http.ResponseWriter, req *http.Request) {
	log.Print("/ip/{id}")
	vars := mux.Vars(req)
	m, err := system.Get1AdresseIP(vars["id"])
	if err != nil {
		log.Printf("Feeler /ip/{id} : %v", err)
	}
	system.EnvoiJSON(m, "/ip/{id}", w)
}

func HTTPcpu(w http.ResponseWriter, _ *http.Request) {
	log.Print("/cpu")
	m, err := system.GetProcesseur()
	if err != nil {
		log.Printf("Feeler /cpu : %v", err)
	}
	system.EnvoiJSON(m, "/cpu", w)
}

func HTTPdisque(w http.ResponseWriter, _ *http.Request) {
	log.Print("/disque")
	m, err := system.GetDisque()
	if err != nil {
		log.Printf("Feeler /disque : %v", err)
	}
	system.EnvoiJSON(m, "/disque", w)
}

func HTTPapi(w http.ResponseWriter, _ *http.Request) {
	log.Print("/api")
	system.EnvoiJSON(system.GetAPI(), "/api", w)
}

func main() {
	log.Print("Ufank vun :8090")
	r := mux.NewRouter()
	r.HandleFunc("/api", HTTPapi)
	r.HandleFunc("/cpu", HTTPcpu)
	r.HandleFunc("/disque", HTTPdisque)
	r.HandleFunc("/ip/passerelle", HTTPippasserelle)
	r.HandleFunc("/ip/{id}", HTTP1ip)
	r.HandleFunc("/ip", HTTPip)
	// Fir ofzeschléissen ...
	http.ListenAndServe(":8090", r)
}
