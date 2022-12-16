package system

import (
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

// Get1CarteReseau : Carte réseau
func Get1CarteReseau(name string) (CarteReseau, error) {
	var c CarteReseau

	cards, err := net.Interfaces()
	if err != nil {
		return CarteReseau{}, err
	}
	infos, err := net.IOCounters(true)
	if err != nil {
		return CarteReseau{}, err
	}

	for i, card := range cards {
		if infos[i].Name == name {
			c.Name = infos[i].Name
			c.MTU = card.MTU
			c.MacAddress = card.HardwareAddr
			c.BytesSent = infos[i].BytesSent
			c.BytesRecv = infos[i].BytesRecv
			c.PacketsSent = infos[i].PacketsSent
			c.PacketsRecv = infos[i].PacketsRecv
			c.ErrorsIn = infos[i].Errin
			c.ErrorsOut = infos[i].Errout
			c.Addrs = card.Addrs
			break
		}
	}
	return c, nil
}

// GetCarteReseau : Cartes réseaux
func GetCarteReseau() ([]CarteReseau, error) {
	var result []CarteReseau

	cards, err := net.Interfaces()
	if err != nil {
		return []CarteReseau{}, err
	}
	infos, err := net.IOCounters(true)
	if err != nil {
		return []CarteReseau{}, err
	}

	for i, card := range cards {
		var c CarteReseau
		c.Name = infos[i].Name
		c.MTU = card.MTU
		c.MacAddress = card.HardwareAddr
		c.BytesSent = infos[i].BytesSent
		c.BytesRecv = infos[i].BytesRecv
		c.PacketsSent = infos[i].PacketsSent
		c.PacketsRecv = infos[i].PacketsRecv
		c.ErrorsIn = infos[i].Errin
		c.ErrorsOut = infos[i].Errout
		c.Addrs = card.Addrs
		result = append(result, c)
	}
	return result, nil
}
