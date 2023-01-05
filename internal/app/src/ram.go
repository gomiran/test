package src

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// RAM details in host
type RAM struct {
	Free  uint64 `json:"free"`
	Total uint64 `json:"total"`
	Usage uint64 `json:"usage"`
}

func CheckRAM() RAM {

	ramChan := make(chan RAM)
	go func(c chan RAM) {
		memory, err := mem.VirtualMemory()
		if err != nil {
			fmt.Print(err)
		}
		c <- RAM{
			Free:  memory.Total/GB - memory.Used/GB,
			Usage: memory.Used / GB,
			Total: memory.Total / GB,
		}
	}(ramChan)

	return <-ramChan
}
