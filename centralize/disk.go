package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Disque struct {
	ID          int     `json:"id"`
	Device      string  `json:"device"`
	MountPoint  string  `json:"mountpoint"`
	FSType      string  `json:"fstype"`
	UsedPerCent float64 `json:"usedpcent"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
}

func Get1PCDisque(ip string) ([]Disque, error) {
	resp, err := http.Get("http://" + ip + ":8090/disque")
	if err != nil {
		return []Disque{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []Disque{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []Disque{}, err
	}

	return ps, nil
}

func Get1PC1Disque(ip string, id string) (Disque, error) {
	resp, err := http.Get("http://" + ip + ":8090/disque/" + id)
	if err != nil {
		return Disque{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := Disque{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return Disque{}, err
	}

	return ps, nil
}
