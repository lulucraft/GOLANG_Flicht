package system

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
)

type AdresseIP struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Adresse string `json:"ip"`
	Netmask string `json:"netmask"`
}

// netmask : transformation du masque ivp4
func netmask(val string) string {
	var tab [4]int64
	tab[0], _ = strconv.ParseInt(val[0:2], 16, 64)
	tab[1], _ = strconv.ParseInt(val[2:4], 16, 64)
	tab[2], _ = strconv.ParseInt(val[4:6], 16, 64)
	tab[3], _ = strconv.ParseInt(val[6:8], 16, 64)

	return fmt.Sprint(tab[0], ".", tab[1], ".", tab[2], ".", tab[3])
}

// netmask6 : transformation du masque ivp6
func netmask6(val string) string {
	var r string

	for i := 0; i < 16; i++ {
		r = r + val[i:i+2] + ":"
	}

	return r[:len(r)-1]
}

// Get1AdresseIP : Récupération d'une adresse IP
func Get1AdresseIP(id string) (AdresseIP, error) {
	n, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return AdresseIP{}, err
	}
	num := int((n))

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return AdresseIP{}, err
	}
	ip := AdresseIP{}
	for i, address := range addrs {
		if i != num {
			continue
		}
		var mask, itype string
		ipnet, _ := address.(*net.IPNet)
		if ipnet.IP.To4() != nil {
			mask = netmask(ipnet.Mask.String())
			itype = "ipv4"

		} else {
			mask = netmask6(ipnet.Mask.String())
			itype = "ipv6"
		}
		ip = AdresseIP{
			ID:      i,
			Netmask: mask,
			Adresse: ipnet.IP.String(),
			Type:    itype,
		}
		break
	}
	return ip, nil
}

// GetAdresseIP : Récupération de l'adresse IP
func GetAdresseIP() ([]AdresseIP, error) {
	ips := []AdresseIP{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return []AdresseIP{}, err
	}
	for i, address := range addrs {
		var mask string
		ipnet, _ := address.(*net.IPNet)
		if ipnet.IP.To4() != nil {
			mask = netmask(ipnet.Mask.String())

		} else {
			mask = netmask6(ipnet.Mask.String())
		}
		ip := AdresseIP{
			ID:      i,
			Netmask: mask,
			Adresse: ipnet.IP.String(),
			Type:    "local",
		}
		ips = append(ips, ip)
	}
	return ips, nil
}

// GetPasserelle : IP Externe
func GetPasserelle() (AdresseIP, error) {
	resp, err := http.Get("http://ifconfig.me")
	if err != nil {
		return AdresseIP{}, err
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	ip := AdresseIP{
		Type:    "passerelle",
		Adresse: string(content),
		Netmask: "0.0.0.0",
	}
	return ip, nil
}
