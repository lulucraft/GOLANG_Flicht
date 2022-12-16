package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AdresseIP struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Adresse string `json:"ip"`
	Netmask string `json:"netmask"`
}

func Get1PCAdresseIP(ip string) ([]AdresseIP, error) {
	resp, err := http.Get("http://" + ip + ":8090/ip")
	if err != nil {
		return []AdresseIP{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []AdresseIP{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []AdresseIP{}, err
	}

	return ps, nil
}

func Get1PC1AdresseIP(ip string, id string) (AdresseIP, error) {
	resp, err := http.Get("http://" + ip + ":8090/ip/" + id)
	if err != nil {
		return AdresseIP{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := AdresseIP{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return AdresseIP{}, err
	}

	return ps, nil
}

func Get1PCPasserelle(ip string) (AdresseIP, error) {
	resp, err := http.Get("http://" + ip + ":8090/passerelle")
	if err != nil {
		return AdresseIP{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := AdresseIP{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return AdresseIP{}, err
	}

	return ps, nil
}
