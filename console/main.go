package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Memoire struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"Available"`
	Used      uint64 `json:"Used"`
	Free      uint64 `json:"Free"`
}

type Processeur struct {
	ID       int32   `json:"id"`
	VendorID string  `json:"vendoid"`
	Family   string  `json:"family"`
	Model    string  `json:"model"`
	Mhz      float64 `json:"freq"`
	PerCent  float64 `json:"pcent"`
}

type CarteReseau struct {
	Name        string `json:"name"`
	MTU         int    `json:"mtu"`
	MacAddress  string `json:"id"`
	BytesSent   uint64 `json:"bytesSent"`
	BytesRecv   uint64 `json:"bytesRecv"`
	PacketsSent uint64 `json:"packetsSent"`
	PacketsRecv uint64 `json:"packetsRecv"`
	ErrorsIn    uint64 `json:"errin"`
	ErrorsOut   uint64 `json:"errout"`
	//Addrs       []net.InterfaceAddr
}
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
type AdresseIP struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Adresse string `json:"ip"`
	Netmask string `json:"netmask"`
}

type Processus struct {
	User   string   `json:"user"`
	Pid    int32    `json:"pid"`
	Cpu    float64  `json:"cpu"`
	Mem    float32  `json:"mem"`
	Name   string   `json:"name"`
	Status []string `json:"status"`
}

type Charge struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func GetMemory(ipaddr string) (Memoire, error) {
	// Recup
	resp, err := http.Get("http://192.168.192.232:8090/api/" + ipaddr + "/memoire")
	if err != nil {
		return Memoire{}, err
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)

	// Parsing
	memoire := Memoire{}
	err = json.Unmarshal(content, &memoire)
	if err != nil {
		return Memoire{}, err
	}

	return memoire, nil
}

func GetCharge(ipaddr string) (Charge, error) {
	// Recup
	resp, err := http.Get("http://192.168.192.232:8090/api/" + ipaddr + "/charge")
	if err != nil {
		return Charge{}, err
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)

	// Parsing
	charge := Charge{}
	err = json.Unmarshal(content, &charge)
	if err != nil {
		return Charge{}, err
	}

	return charge, nil

}

func format(val uint64) string {
	f := float64(val) / 1024 / 1024 / 1024
	return fmt.Sprintf("%.2f Go", f)

}

func main() {
	var ipadd string
	var choix int
	fmt.Println("choix de la machine")
	fmt.Println("rentrez l'adresse ip :")
	fmt.Scan(&ipadd)
	fmt.Println("Que voulez-vous voir ?")
	fmt.Println("1 : Memoire  ou 2 : Charge ")
	fmt.Scan(&choix)

	for {
		if choix == 1 {
			m, err := GetMemory(ipadd)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Mémoire : %v / %v\n", format(m.Used), format(m.Total))
			time.Sleep(2 * time.Second)
		} else {
			m, err := GetCharge(ipadd)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Charges : charge 1 = %v charge 5 = %v charge 15 = %v \n", m.Load1, m.Load5, m.Load15)
			time.Sleep(2 * time.Second)
		}

	}

}
