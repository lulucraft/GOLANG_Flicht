package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Processus struct {
	User   string   `json:"user"`
	Pid    int32    `json:"pid"`
	Cpu    float64  `json:"cpu"`
	Mem    float32  `json:"mem"`
	Name   string   `json:"name"`
	Status []string `json:"status"`
}

func Get1PCProcessus(ip string) ([]Processus, error) {
	resp, err := http.Get("http://" + ip + ":8090/processus")
	if err != nil {
		return []Processus{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []Processus{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []Processus{}, err
	}

	return ps, nil
}

func Get1PC1Processus(ip string, id string) (Processus, error) {
	resp, err := http.Get("http://" + ip + ":8090/processus/" + id)
	if err != nil {
		return Processus{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := Processus{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return Processus{}, err
	}

	return ps, nil
}
