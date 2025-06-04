package system

import "github.com/StackExchange/wmi"

type Win32_Processor struct {
	Name                      string
	Manufacturer              string
	NumberOfCores             uint32
	NumberOfLogicalProcessors uint32
	MaxClockSpeed             uint32
}

type CPUInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Cores        uint32 `json:"cores"`
	Threads      uint32 `json:"threads"`
	Clock        uint32 `json:"clock"`        // MHz
	CurruntClock uint32 `json:"CurruntClock"` // MHz
}

func GetCPUInfo() (*CPUInfo, error) {
	var cpuData []Win32_Processor
	err := wmi.Query("SELECT Name, Manufacturer, NumberOfCores, NumberOfLogicalProcessors, MaxClockSpeed FROM Win32_Processor", &cpuData)
	if err != nil || len(cpuData) == 0 {
		return nil, err
	}

	c := cpuData[0]
	return &CPUInfo{
		Manufacturer: c.Manufacturer,
		Model:        c.Name,
		Cores:        c.NumberOfCores,
		Threads:      c.NumberOfLogicalProcessors,
		Clock:        c.MaxClockSpeed,
	}, nil
}
