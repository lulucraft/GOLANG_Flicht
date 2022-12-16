package system

import (
	"github.com/shirou/gopsutil/v3/load"
)

type Charge struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

// GetCharge : Récupération de la charge
func GetCharge() (Charge, error) {
	info, err := load.Avg()
	if err != nil {
		return Charge{}, err
	}
	l := Charge{
		Load1:  info.Load1,
		Load5:  info.Load5,
		Load15: info.Load15,
	}
	return l, nil
}
