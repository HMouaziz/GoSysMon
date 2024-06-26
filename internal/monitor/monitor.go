package monitor

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func GetMemoryUsage() (uint64, uint64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return v.Used, v.Total, nil
}

func DisplayUsage(interval time.Duration) {
	for {
		cpuUsage, err := GetCPUUsage()
		if err != nil {
			fmt.Println("Error getting CPU usage:", err)
			return
		}

		usedMem, totalMem, err := GetMemoryUsage()
		if err != nil {
			fmt.Println("Error getting memory usage:", err)
			return
		}

		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		fmt.Printf("Memory Usage: %v / %v MB\n", usedMem/1024/1024, totalMem/1024/1024)

		time.Sleep(interval)
	}
}
