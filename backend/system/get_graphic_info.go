package system

import "github.com/StackExchange/wmi"

/* ────────── WMI 원본 구조체 ────────── */

type Win32_VideoController struct {
	Name                 string // 예: "NVIDIA GeForce RTX 4060 Laptop GPU"
	AdapterCompatibility string // 예: "NVIDIA"
	AdapterRAM           uint32 // Bytes (4 GB 초과 시 32-bit 한계!)
	DriverVersion        string // 예: "31.0.15.4624"
}

/* ────────── 앱에서 쓰는 구조체 ────────── */

type GPUInfo struct {
	Manufacturer  string `json:"manufacturer,omitempty"`
	Model         string `json:"model,omitempty"`
	VRAMMB        uint64 `json:"vramMB"` // MB
	DriverVersion string `json:"driverVersion,omitempty"`
	ClockMHz      uint32 `json:"clock,omitempty"` // vendor API 필요
}

type GPUs struct {
	Cards []GPUInfo `json:"cards"`
	Count int       `json:"count"`
}

/* ────────── GPU 정보 조회 함수 ────────── */

func GetGPUInfo() (*GPUs, error) {
	var raw []Win32_VideoController
	q := `SELECT Name, AdapterCompatibility, AdapterRAM, DriverVersion
	      FROM Win32_VideoController`
	if err := wmi.Query(q, &raw); err != nil || len(raw) == 0 {
		return nil, err
	}

	gpus := &GPUs{}
	for _, v := range raw {
		g := GPUInfo{
			Manufacturer:  v.AdapterCompatibility,
			Model:         v.Name,
			VRAMMB:        uint64(v.AdapterRAM) / (1024 * 1024), // Bytes → MB
			DriverVersion: v.DriverVersion,
			ClockMHz:      0, // WMI 미제공. 필요 시 NVAPI/ADL 등으로 보강
		}
		gpus.Cards = append(gpus.Cards, g)
	}
	gpus.Count = len(gpus.Cards)
	return gpus, nil
}
