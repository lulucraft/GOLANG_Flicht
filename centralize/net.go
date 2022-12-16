package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shirou/gopsutil/v3/net"
)

type CarteReseau struct {
	Name        string
	MTU         int
	MacAddress  string
	BytesSent   uint64
	BytesRecv   uint64
	PacketsSent uint64
	PacketsRecv uint64
	ErrorsIn    uint64
	ErrorsOut   uint64
	Addrs       []net.InterfaceAddr
}

func Get1PC1CarteReseau(ip string, name string) (CarteReseau, error) {
	resp, err := http.Get("http://" + ip + ":8090/carte/" + name)
	if err != nil {
		return CarteReseau{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := CarteReseau{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return CarteReseau{}, err
	}

	return ps, nil
}

func Get1PCCarteReseau(ip string) ([]CarteReseau, error) {
	resp, err := http.Get("http://" + ip + ":8090/carte")
	if err != nil {
		return []CarteReseau{}, err
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	ps := []CarteReseau{}
	err = json.Unmarshal(content, &ps)
	if err != nil {
		return []CarteReseau{}, err
	}

	return ps, nil
}
