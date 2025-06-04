package system

import (
	"github.com/StackExchange/wmi"
)

/* ────────── WMI 원본 구조체 ────────── */

type Win32_PhysicalMemory struct {
	Manufacturer     string
	PartNumber       string
	Capacity         uint64 // Bytes
	Speed            uint32 // MHz
	MemoryType       uint16
	SMBIOSMemoryType uint8
}

/* ────────── 앱에서 사용하는 구조체 ────────── */

type RAMModule struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"` // PartNumber
	CapacityMB   uint64 `json:"capacityMB"`      // MB
	SpeedMHz     uint32 `json:"speedMHz,omitempty"`
	DDR          string `json:"ddr,omitempty"` // DDR3·DDR4·DDR5 …
}

type RAMInfo struct {
	Modules []RAMModule `json:"modules"`
	TotalMB uint64      `json:"totalMB"`
	Count   int         `json:"count"`
}

/* ────────── DDR 매핑 & 휴리스틱 ────────── */

// SMBIOS 값(DDR3~5)
var smbiosDDR = map[uint8]string{
	0x1A: "DDR3", // 26
	0x1B: "DDR4", // 27
	0x1C: "DDR5", // 28
}

// 구형 MemoryType 매핑
var memTypeDDR = map[uint16]string{
	21: "DDR2",
	22: "DDR2 FB-DIMM",
	24: "DDR",
	26: "DDR4", // 일부 BIOS
	27: "DDR5",
}

// 속도로 추정 (단위: MHz)
func guessDDRFromSpeed(mhz uint32) string {
	switch {
	case mhz >= 4800:
		return "DDR5"
	case mhz >= 2133:
		return "DDR4"
	case mhz >= 1066:
		return "DDR3"
	case mhz >= 533:
		return "DDR2"
	}
	return ""
}

func ddrType(memType uint16, smbios uint8, speed uint32) string {
	if v, ok := smbiosDDR[smbios]; ok {
		return v
	}
	if v, ok := memTypeDDR[memType]; ok {
		return v
	}
	return guessDDRFromSpeed(speed)
}

/* ────────── RAM 정보 조회 함수 ────────── */

func GetRAMInfo() (*RAMInfo, error) {
	var raw []Win32_PhysicalMemory
	q := `SELECT Manufacturer, PartNumber, Capacity, Speed,
	      MemoryType, SMBIOSMemoryType FROM Win32_PhysicalMemory`
	if err := wmi.Query(q, &raw); err != nil || len(raw) == 0 {
		return nil, err
	}

	info := &RAMInfo{}
	for _, m := range raw {
		mod := RAMModule{
			Manufacturer: m.Manufacturer,
			Model:        m.PartNumber,
			CapacityMB:   m.Capacity / (1024 * 1024), // Bytes → MB
			SpeedMHz:     m.Speed,
			DDR:          ddrType(m.MemoryType, m.SMBIOSMemoryType, m.Speed),
		}
		info.Modules = append(info.Modules, mod)
		info.TotalMB += mod.CapacityMB
	}
	info.Count = len(info.Modules)
	return info, nil
}
