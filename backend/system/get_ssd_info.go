package system

import (
	"strings"

	"github.com/StackExchange/wmi"
)

/* ────────── WMI 구조체 ────────── */

// 기본 드라이브 정보 (CIMV2)
type Win32_DiskDrive struct {
	Model         string
	Manufacturer  string
	SerialNumber  string
	InterfaceType string
	Size          uint64 // Bytes
	MediaType     string // 보통 "Fixed hard disk media"
}

// SSD/HDD 구분용 (Windows 8+)
type MSFT_PhysicalDisk struct {
	SerialNumber string
	MediaType    uint16 // 3=HDD, 4=SSD, 5=SCM
}

/* ────────── 앱에서 쓰는 구조체 ────────── */

type DiskInfo struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
	Serial       string `json:"serial,omitempty"`
	Interface    string `json:"interface,omitempty"` // SATA 대신 IDE·SCSI 등으로 표기
	SizeGB       uint64 `json:"sizeGB"`
	Type         string `json:"type,omitempty"` // SSD·HDD·SCM
}

type Disks struct {
	Drives []DiskInfo `json:"drives"`
	Count  int        `json:"count"`
}

/* ────────── MediaType 매핑 ────────── */

func mediaEnumToStr(code uint16) string {
	switch code {
	case 3:
		return "HDD"
	case 4:
		return "SSD"
	case 5:
		return "SCM"
	default:
		return ""
	}
}

// ROOT\Microsoft\Windows\Storage 쿼리 → Serial→Type 맵
func loadMediaTypeMap() map[string]string {
	var phys []MSFT_PhysicalDisk
	err := wmi.QueryNamespace(
		`SELECT SerialNumber, MediaType FROM MSFT_PhysicalDisk`,
		&phys, `ROOT\Microsoft\Windows\Storage`)
	if err != nil {
		return nil // Windows 7이거나 WMI Provider가 없을 때
	}

	mt := make(map[string]string, len(phys))
	for _, p := range phys {
		if s := strings.TrimSpace(p.SerialNumber); s != "" {
			mt[s] = mediaEnumToStr(p.MediaType)
		}
	}
	return mt
}

/* ────────── 최종 함수 ────────── */

func GetDiskInfo() (*Disks, error) {
	// 1) Win32_DiskDrive: 기본 스펙
	var raw []Win32_DiskDrive
	if err := wmi.Query(
		`SELECT Model, Manufacturer, SerialNumber, InterfaceType, Size, MediaType FROM Win32_DiskDrive`,
		&raw); err != nil || len(raw) == 0 {
		return nil, err
	}

	// 2) (가능하면) MediaType 맵 획득
	mtMap := loadMediaTypeMap()

	out := &Disks{}
	for _, d := range raw {
		typ := "" // SSD/HDD/SCM 판정

		// ① W10+ MediaType 열 우선
		if mtMap != nil {
			if v, ok := mtMap[strings.TrimSpace(d.SerialNumber)]; ok {
				typ = v
			}
		}
		// ② WMI 텍스트·모델명 휴리스틱
		if typ == "" {
			lm := strings.ToLower(d.MediaType + d.Model)
			switch {
			case strings.Contains(lm, "ssd"), strings.Contains(lm, "solid"):
				typ = "SSD"
			case strings.Contains(lm, "hdd"), strings.Contains(lm, "hard"):
				typ = "HDD"
			}
		}

		out.Drives = append(out.Drives, DiskInfo{
			Manufacturer: d.Manufacturer,
			Model:        d.Model,
			Serial:       strings.TrimSpace(d.SerialNumber),
			Interface:    d.InterfaceType,
			SizeGB:       d.Size / (1024 * 1024 * 1024),
			Type:         typ,
		})
	}
	out.Count = len(out.Drives)
	return out, nil
}
