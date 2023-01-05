package app

const (
	HostInfo         = "hostInfo"
	Cpu              = "cpu"
	Ram              = "ram"
	Disks            = "disks"
	NetworkDevices   = "networkDevices"
	NetworkBandwidth = "networkBandwidth"
	Processes        = "processes"
)

var mapModules = map[string]bool{
	HostInfo:         true,
	Cpu:              true,
	Ram:              true,
	Disks:            true,
	NetworkDevices:   true,
	NetworkBandwidth: true,
	Processes:        true,
}

func available(module string) bool {
	return mapModules[module]
}
