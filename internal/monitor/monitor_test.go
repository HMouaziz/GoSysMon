package monitor

import (
	"testing"
)

func TestGetCPUUsage(t *testing.T) {
	usage, err := GetCPUUsage()
	if err != nil {
		t.Errorf("Error getting CPU usage: %v", err)
	}
	if usage < 0 {
		t.Errorf("CPU usage should not be negative: %v", usage)
	}
}

func TestGetMemoryUsage(t *testing.T) {
	used, total, err := GetMemoryUsage()
	if err != nil {
		t.Errorf("Error getting memory usage: %v", err)
	}
	if used > total {
		t.Errorf("Used memory should not exceed total memory: %v / %v", used, total)
	}
}
