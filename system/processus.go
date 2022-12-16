package system

import (
	"strconv"

	"github.com/shirou/gopsutil/v3/process"
)

type Processus struct {
	User   string   `json:"user"`
	Pid    int32    `json:"pid"`
	Cpu    float64  `json:"cpu"`
	Mem    float32  `json:"mem"`
	Name   string   `json:"name"`
	Status []string `json:"status"`
}

// Get1Processus : récupération d'un processus
func Get1Processus(id string) (Processus, error) {
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return Processus{}, err
	}

	processes, err := process.Processes()
	if err != nil {
		return Processus{}, err
	}
	var ps Processus
	for _, process := range processes {
		if int32(i) == process.Pid {
			ps.User, _ = process.Username()
			ps.Pid = process.Pid
			ps.Cpu, _ = process.CPUPercent()
			ps.Mem, _ = process.MemoryPercent()
			ps.Name, _ = process.Name()
			ps.Status, _ = process.Status()
			break
		}
	}
	return ps, nil
}

// GetProcessus : récupération des processus
func GetProcessus() ([]Processus, error) {
	processes, err := process.Processes()
	if err != nil {
		return []Processus{}, err
	}
	var pslist []Processus
	for _, process := range processes {
		var ps Processus
		ps.User, _ = process.Username()
		ps.Pid = process.Pid
		ps.Cpu, _ = process.CPUPercent()
		ps.Mem, _ = process.MemoryPercent()
		ps.Name, _ = process.Name()
		ps.Status, _ = process.Status()
		pslist = append(pslist, ps)
	}
	return pslist, nil
}

func KillProcessus(id string) (string, error) {
	pid, err := strconv.ParseInt(id, 10, 32)

	processes, err := process.Processes()
	if err != nil {
		return "no process found", err
	}
	for _, p := range processes {
		i := p.Pid
		if i == int32(pid) {
			err := p.Kill()
			if err != nil {
				return "kill error", err
			}
			return "kill ok", nil
		}
	}
	return "process not found", nil
}
