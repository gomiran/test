package main

import (
	"fmt"
	"monitorwin/internal/app"
	"monitorwin/internal/app/src"
	"net"
	"net/http"
	"strconv"
)

var availableConfigPort string

func portAvailable(configPort string) {
	fmt.Println("Setting port to " + configPort)
	ln, err := net.Listen("tcp", ":"+configPort)
	if err != nil {
		fmt.Println("Can't listen on port " + configPort)
		newConfigPort, errStrConv := strconv.Atoi(configPort)
		if errStrConv != nil {
			panic(errStrConv)
		}

		newConfigPort = newConfigPort + 1
		newConfigPortString := strconv.Itoa(newConfigPort)
		portAvailable(newConfigPortString)
	} else {
		_ = ln.Close()
		availableConfigPort = configPort
	}
}

func main() {

	portAvailable("4000")

	fmt.Println("Starting Server at port " + availableConfigPort)
	app.CreateEndPoint("hostInfo", src.CheckHostInfo())
	err := http.ListenAndServe(":"+availableConfigPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
