package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Charge struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func Get1PCCharge(ip string) (Charge, error) {
	resp, err := http.Get("http://" + ip + ":8090/charge")
	if err != nil {
		return Charge{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := Charge{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return Charge{}, err
	}

	return ps, nil
}
