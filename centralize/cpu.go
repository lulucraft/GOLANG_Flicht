package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Processeur struct {
	ID       int32   `json:"id"`
	VendorID string  `json:"vendoid"`
	Family   string  `json:"family"`
	Model    string  `json:"model"`
	Mhz      float64 `json:"freq"`
	PerCent  float64 `json:"pcent"`
}

func Get1PCProcesseur(ip string) ([]Processeur, error) {
	resp, err := http.Get("http://" + ip + ":8090/processeur")
	if err != nil {
		return []Processeur{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []Processeur{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []Processeur{}, err
	}

	return ps, nil
}
