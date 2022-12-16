package system

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type API struct {
	URL     string `json:"url"`
	Methode string `json:"methode"`
}

// GetAPI : Liste des méthodes par API
func GetAPI() []API {
	return []API{
		{
			URL:     "/api",
			Methode: "GetAPI",
		},
		{
			URL:     "/processeur",
			Methode: "GetProcesseur",
		},
		{
			URL:     "/disque",
			Methode: "GetDisque",
		},
		{
			URL:     "/disque/{id}",
			Methode: "Get1Disque",
		},
		{
			URL:     "/ip",
			Methode: "GetAdresseIP",
		},
		{
			URL:     "/ip/{id}",
			Methode: "Get1AdresseIP",
		},
		{
			URL:     "/ip/passerelle",
			Methode: "GetPasserelle",
		},
		{
			URL:     "/charge",
			Methode: "GetCharge",
		},
		{
			URL:     "/memoire",
			Methode: "GetMemoire",
		},
		{
			URL:     "/carte",
			Methode: "GetCarteReseau",
		},
		{
			URL:     "/carte/{name}",
			Methode: "Get1CarteReseau",
		},
		{
			URL:     "/processus",
			Methode: "GetProcessus",
		},
		{
			URL:     "/processus/{id}",
			Methode: "Get1Processus",
		},
		{
			URL:     "/processus/kill/{id}",
			Methode: "Kill1Processus",
		},
	}
}

// EnvoiJSON : envoie le JSON
func EnvoiJSON(ps any, url string, w http.ResponseWriter) {
	// Transformation JSON
	j, err := json.Marshal(ps)
	if err != nil {
		log.Printf("ERREUR %s : %v", url, err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))
}
