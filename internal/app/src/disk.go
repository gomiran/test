package src

import "github.com/shirou/gopsutil/disk"

// multipli sizes by 1024 to get bytes
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PT = 1024 * TB
)

type Sizes struct {
	KB float64
	MB float64
	GB float64
	TB float64
	PT float64
}

// func kb uint to gb uint
func KbToGb(kb uint64) uint64 {
	return kb / 1024 / 1024
}

// return list of sizes
func (d Disk) SizeFree() Sizes {
	return Sizes{
		KB: float64(d.Free) / KB,
		MB: float64(d.Free) / MB,
		GB: float64(d.Free) / GB,
		TB: float64(d.Free) / TB,
		PT: float64(d.Free) / PT,
	}
}
func (d Disk) SizeUsed() Sizes {
	return Sizes{
		KB: float64(d.Used) / KB,
		MB: float64(d.Used) / MB,
		GB: float64(d.Used) / GB,
		TB: float64(d.Used) / TB,
		PT: float64(d.Used) / PT,
	}
}
func (d Disk) SizeSize() Sizes {
	return Sizes{
		KB: float64(d.Size) / KB,
		MB: float64(d.Size) / MB,
		GB: float64(d.Size) / GB,
		TB: float64(d.Size) / TB,
		PT: float64(d.Size) / PT,
	}
}

// Disk properties
type Disk struct {
	Mountpoint string  `json:"mountPoint"`
	Free       uint64  `json:"free"`
	Size       uint64  `json:"size"`
	Used       uint64  `json:"used"`
	Percent    float64 `json:"percent"`
}

type Disks []Disk

func CheckDisks() Disks {

	disksChan := make(chan Disks)
	go func(c chan Disks) {
		disks, _ := disk.Partitions(false)

		var totalDisks Disks

		diskChan := make(chan Disk)
		for _, d := range disks {
			go func(d disk.PartitionStat) {
				diskUsageOf, _ := disk.Usage(d.Mountpoint)
				diskChan <- Disk{
					Free:       diskUsageOf.Free / GB,
					Mountpoint: d.Mountpoint,
					Percent:    diskUsageOf.UsedPercent / GB,
					Size:       diskUsageOf.Total / GB,
					Used:       diskUsageOf.Used / GB,
				}

			}(d)

		}
		taskList := len(disks)

		for d := range diskChan {
			totalDisks = append(totalDisks, d)
			taskList--
			if taskList == 0 {
				break
			}
		}
		c <- totalDisks
	}(disksChan)

	return <-disksChan
}
