package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Memoire struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
}

func Get1PCMemoire(ip string) (Memoire, error) {
	resp, err := http.Get("http://" + ip + ":8090/memoire")
	if err != nil {
		return Memoire{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := Memoire{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return Memoire{}, err
	}

	return ps, nil
}
