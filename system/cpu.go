package system

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type Processeur struct {
	ID       int32   `json:"id"`
	VendorID string  `json:"vendoid"`
	Family   string  `json:"family"`
	Model    string  `json:"model"`
	Mhz      float64 `json:"freq"`
	PerCent  float64 `json:"pcent"`
}

// GetProcesseur : Récupération des processeurs
func GetProcesseur() ([]Processeur, error) {

	var cpus []Processeur

	cpuInfos, err := cpu.Info()
	if err != nil {
		return []Processeur{}, err
	}

	percents, err := cpu.Percent(time.Second, true)
	if err != nil {
		return []Processeur{}, err
	}

	i := 0
	for _, ci := range cpuInfos {
		cpu := Processeur{}
		cpu.ID = ci.CPU
		cpu.VendorID = ci.VendorID
		cpu.Family = ci.Family
		cpu.Model = ci.Model
		cpu.Mhz = ci.Mhz
		cpu.PerCent = percents[i]
		cpus = append(cpus, cpu)
		i++
		// cpus = append(cpus, cpu)
	}

	return cpus, nil
}
