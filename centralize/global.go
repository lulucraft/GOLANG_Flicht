package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type API struct {
	URL     string `json:"url"`
	Methode string `json:"methode"`
}

func Get1PCApi(ip string) ([]API, error) {
	resp, err := http.Get("http://" + ip + ":8090/api")
	if err != nil {
		return []API{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []API{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []API{}, err
	}

	return ps, nil
}
