package main

import (
	"log"
	"net/http"

	"git.evoliatis.fr/scollado/system"
	"github.com/gorilla/mux"
)

// /api
func HTTPapi(w http.ResponseWriter, _ *http.Request) {
	log.Print("/api")
	system.EnvoiJSON(system.GetAPI(), "/api", w)
}

func HTTPapi1PCapi(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/api")
	vars := mux.Vars(req)
	m, err := Get1PCApi(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/api : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/api", w)
}

// /processeur
func HTTPapi1PCprocesseur(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/processeur")
	vars := mux.Vars(req)
	m, err := Get1PCProcesseur(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/processeur : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/processeur", w)
}

// /carte
func HTTPapi1PCcarte(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/carte")
	vars := mux.Vars(req)
	m, err := Get1PCCarteReseau(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/carte : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/carte", w)
}

func HTTPapi1PC1carte(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/carte/{id}")
	vars := mux.Vars(req)
	m, err := Get1PC1CarteReseau(vars["ip"], vars["id"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/carte/{id} : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/carte/{id}", w)
}

// /charge
func HTTPapi1PCcharge(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/charge")
	vars := mux.Vars(req)
	m, err := Get1PCCharge(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/charge : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/charge", w)
}

// /memoire
func HTTPapi1PCmemoire(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/memoire")
	vars := mux.Vars(req)
	m, err := Get1PCMemoire(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/memoire : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/memoire", w)
}

// /ip
func HTTPapi1PCip(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/ip")
	vars := mux.Vars(req)
	m, err := Get1PCAdresseIP(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/ip : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/ip", w)
}

func HTTPapi1PC1ip(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/ip/{id}")
	vars := mux.Vars(req)
	m, err := Get1PC1AdresseIP(vars["ip"], vars["id"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/ip/{id} : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/ip/{id}", w)
}

func HTTPapi1PCippasserelle(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/ip/passerelle")
	vars := mux.Vars(req)
	m, err := Get1PCPasserelle(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/ip/passerelle : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/ip/passerelle", w)
}

// /processus
func HTTPapi1PCprocessus(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/processus")
	vars := mux.Vars(req)
	m, err := Get1PCProcessus(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/processus : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/processus", w)
}

func HTTPapi1PC1processus(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/processus/{id}")
	vars := mux.Vars(req)
	m, err := Get1PC1Processus(vars["ip"], vars["id"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/processus/{id} : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/processus/{id}", w)
}

// /disque
func HTTPapi1PCdisque(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/disque")
	vars := mux.Vars(req)
	m, err := Get1PCDisque(vars["ip"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/disque : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/disque", w)
}

func HTTPapi1PC1disque(w http.ResponseWriter, req *http.Request) {
	log.Print("/api/{ip}/disque/{id}")
	vars := mux.Vars(req)
	m, err := Get1PC1Disque(vars["ip"], vars["id"])
	if err != nil {
		log.Printf("Erreur /api/{ip}/disque/{id} : %v", err)
	}
	system.EnvoiJSON(m, "/api/{ip}/disque/{id}", w)
}

// main
func main() {
	log.Print("Listen on port 8090")

	r := mux.NewRouter()

	// Api
	r.HandleFunc("/api", HTTPapi)
	r.HandleFunc("/api/{ip}/api", HTTPapi1PCapi)

	// Cpu
	r.HandleFunc("/api/{ip}/processeur", HTTPapi1PCprocesseur)

	// Carte
	r.HandleFunc("/api/{ip}/carte", HTTPapi1PCcarte)
	r.HandleFunc("/api/{ip}/carte/{id}", HTTPapi1PC1carte)

	// Charge
	r.HandleFunc("/api/{ip}/charge", HTTPapi1PCcharge)

	// Memoire
	r.HandleFunc("/api/{ip}/memoire", HTTPapi1PCmemoire)

	// Ip
	r.HandleFunc("/api/{ip}/ip", HTTPapi1PCip)
	r.HandleFunc("/api/{ip}/ip/{id}", HTTPapi1PC1ip)
	r.HandleFunc("/api/{ip}/ip/passerelle", HTTPapi1PCippasserelle)

	// Processus
	r.HandleFunc("/api/{ip}/processus", HTTPapi1PCprocessus)
	r.HandleFunc("/api/{ip}/processus/{id}", HTTPapi1PC1processus)

	// Disque
	r.HandleFunc("/api/{ip}/disque", HTTPapi1PCdisque)
	r.HandleFunc("/api/{ip}/disque/{id}", HTTPapi1PC1disque)

	http.ListenAndServe(":8090", r)
}
