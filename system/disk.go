package system

import (
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
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

// GetDisque : Récupération des disques
func GetDisque() ([]Disque, error) {

	parts, err := disk.Partitions(true)
	if err != nil {
		return []Disque{}, err
	}

	var disks []Disque
	for i, part := range parts {
		var d Disque
		t := strings.Split(part.Device, "/")
		if t[0] == "" {
			diskInfo, _ := disk.Usage(part.Mountpoint)
			d.ID = i
			d.Device = part.Device
			d.MountPoint = part.Mountpoint
			d.FSType = part.Fstype
			d.Total = diskInfo.Total
			d.Used = diskInfo.Used
			d.Free = diskInfo.Free
			d.UsedPerCent = diskInfo.UsedPercent
			disks = append(disks, d)
		}
	}

	return disks, nil
}

// Get1Disque : Récupération d'un disque
func Get1Disque(id string) (Disque, error) {
	n, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return Disque{}, err
	}
	num := int((n))

	parts, err := disk.Partitions(true)
	if err != nil {
		return Disque{}, err
	}

	var d Disque
	for i, part := range parts {
		t := strings.Split(part.Device, "/")
		if t[0] == "" && i == num {
			diskInfo, _ := disk.Usage(part.Mountpoint)
			d.ID = i
			d.Device = part.Device
			d.MountPoint = part.Mountpoint
			d.FSType = part.Fstype
			d.Total = diskInfo.Total
			d.Used = diskInfo.Used
			d.Free = diskInfo.Free
			d.UsedPerCent = diskInfo.UsedPercent
			break
		}
	}

	return d, nil
}
