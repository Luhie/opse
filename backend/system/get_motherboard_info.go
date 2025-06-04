package system

import "github.com/StackExchange/wmi"

// WMI 원본 구조체 ─ 필요한 필드만 선언
type Win32_BaseBoard struct {
	Manufacturer string
	Product      string
	SerialNumber string
	Version      string
}

// 앱에서 쓸 가공 구조체
type MotherboardInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Serial       string `json:"serial,omitempty"`
	Version      string `json:"version,omitempty"`
}

// 메인보드 정보 조회
func GetMotherboardInfo() (*MotherboardInfo, error) {
	var boards []Win32_BaseBoard
	q := `SELECT Manufacturer, Product, SerialNumber, Version FROM Win32_BaseBoard`
	if err := wmi.Query(q, &boards); err != nil || len(boards) == 0 {
		return nil, err
	}

	b := boards[0]
	return &MotherboardInfo{
		Manufacturer: b.Manufacturer,
		Model:        b.Product,      // 일부 제조사는 Product 대신 Name 사용
		Serial:       b.SerialNumber, // OEM 보드면 "To be filled by O.E.M." 라고 올 수도 있음
		Version:      b.Version,
	}, nil
}
