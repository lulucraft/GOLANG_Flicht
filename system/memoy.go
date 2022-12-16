package system

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type Memoire struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
}

// GetMemoire : Récupération de la mémoire
func GetMemoire() (Memoire, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return Memoire{}, err
	}
	var mem Memoire
	mem.Available = memInfo.Available
	mem.Free = memInfo.Free
	mem.Total = memInfo.Total
	mem.Used = memInfo.Used

	return mem, nil
}
