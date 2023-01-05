package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"monitorwin/internal/app/src"
)

type data struct {
	HostInfo         src.HostInfo                 `json:"hostInfo"`
	CPU              src.CPU                      `json:"cpu"`
	RAM              src.RAM                      `json:"ram"`
	Disks            []src.Disk                   `json:"disks"`
	NetworkDevices   []src.NetworkDevice          `json:"networkDevices"`
	NetworkBandwidth []src.NetworkDeviceBandwidth `json:"networkBandwidth"`
	Processes        []src.Process                `json:"processes"`
}

func ignoreFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func checkData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("checkData")
	data := data{
		HostInfo:         src.CheckHostInfo(),
		RAM:              src.CheckRAM(),
		CPU:              src.CheckCPU(),
		NetworkDevices:   src.CheckNetworkDevices(),
		NetworkBandwidth: src.CheckNetworkBandwidth(),
		Disks:            src.CheckDisks(),
		Processes:        src.CheckProcesses(),
	}

	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("err")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CreateEndPoint(module string, e interface{}) {
	if available(module) {
		fmt.Println("module: ", module)
		endPoint := func(w http.ResponseWriter, r *http.Request) {
			enableCors(&w)
			data := e
			js, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}

		http.HandleFunc("/"+module, endPoint)
	} else {
		endPoint := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500 - set " + module + " to 'true' in config.yml file"))
		}
		http.HandleFunc("/"+module, endPoint)
	}
}

func init() {
	http.HandleFunc("/", checkData)
	http.HandleFunc("/favicon.ico", ignoreFavicon)
	http.HandleFunc("/host", hostIndex)
	http.HandleFunc("/cpu", cpuIndex)
	http.HandleFunc("/ram", ramIndex)
	http.HandleFunc("/disks", disksIndex)
	http.HandleFunc("/networks", networkIndex)
	http.HandleFunc("/bandwidth", bandwidthIndex)
	http.HandleFunc("/processes", processesIndex)
}

func moduleServer(w http.ResponseWriter, checker interface{}, module string) {
	enableCors(&w)
	if available(module) {
		data := checker
		js, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 - set " + module + " to 'true' in config.yml file"))
	}
}

func hostIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hostIndex")
	moduleServer(w, src.CheckHostInfo(), "hostInfo")
}

func ramIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ramIndex")
	moduleServer(w, src.CheckRAM(), "ram")
}
func cpuIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cpuIndex")

	moduleServer(w, src.CheckCPU(), "cpu")
}

func disksIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("disks")

	moduleServer(w, src.CheckDisks(), "disks")
}

func networkIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, src.CheckNetworkDevices(), "networkDevices")
}

func bandwidthIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("networkBandwidth")

	moduleServer(w, src.CheckNetworkBandwidth(), "networkBandwidth")
}

func processesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processes")

	moduleServer(w, src.CheckProcesses(), "processes")
}
