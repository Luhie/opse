package system

import (
	"os"
	"strconv"
	"strings"

	"github.com/StackExchange/wmi"
)

/* ────────── WMI 구조체 선언 ────────── */

type Win32_NetworkAdapter struct {
	DeviceID        string
	Name            string
	MACAddress      string
	Manufacturer    string
	PhysicalAdapter bool
	AdapterTypeID   *uint16 // 9 = 802.11(Wi-Fi)
}

type Win32_NetworkAdapterConfiguration struct {
	Index     uint32
	IPAddress *[]string
}

/* ────────── 결과 구조체 ────────── */

type NICInfo struct {
	Name   string   `json:"name"`
	MAC    string   `json:"mac"`
	IPs    []string `json:"ips,omitempty"`
	IsWiFi bool     `json:"wifi"`
}

type NetworkInfo struct {
	Hostname string    `json:"hostname"`
	Adapters []NICInfo `json:"adapters"`
}

/* ────────── 내부 헬퍼 ────────── */

func isVirtual(name, mfr string) bool {
	l := strings.ToLower
	vendor := l(mfr)
	desc := l(name)

	switch {
	case strings.Contains(desc, "virtual"):
		return true
	case strings.Contains(desc, "hyper-v"), strings.Contains(vendor, "hyper-v"):
		return true
	case strings.Contains(desc, "vmware"), strings.Contains(vendor, "vmware"):
		return true
	case strings.Contains(desc, "virtualbox"), strings.Contains(vendor, "oracle"):
		return true
	}
	return false
}

/* ────────── 메인 함수 ────────── */

func GetNetworkInfo() (*NetworkInfo, error) {
	// 1) 물리 NIC 목록 가져오기
	var nicsRaw []Win32_NetworkAdapter
	qNic := `SELECT DeviceID, Name, MACAddress, Manufacturer,
	         PhysicalAdapter, AdapterTypeID FROM Win32_NetworkAdapter
	         WHERE PhysicalAdapter = True AND MACAddress IS NOT NULL`
	if err := wmi.Query(qNic, &nicsRaw); err != nil {
		return nil, err
	}

	adapters := make(map[string]*NICInfo) // DeviceID → info
	for _, n := range nicsRaw {
		if isVirtual(n.Name, n.Manufacturer) {
			continue // 가상 어댑터 skip
		}
		id := strings.TrimSpace(n.DeviceID) // Win32_NAC.DeviceID == Index
		adapters[id] = &NICInfo{
			Name:   n.Name,
			MAC:    n.MACAddress,
			IsWiFi: n.AdapterTypeID != nil && *n.AdapterTypeID == 9,
		}
	}

	// 2) IP 주소 붙이기 (IPEnabled=TRUE 만 조회)
	var cfgRaw []Win32_NetworkAdapterConfiguration
	qCfg := `SELECT Index, IPAddress FROM Win32_NetworkAdapterConfiguration
	         WHERE IPEnabled = True`
	_ = wmi.Query(qCfg, &cfgRaw) // IP 없으면 에러 무시

	for _, c := range cfgRaw {
		key := strings.TrimSpace((func(i uint32) string { return strconv.Itoa(int(i)) })(c.Index))
		if nic, ok := adapters[key]; ok && c.IPAddress != nil {
			nic.IPs = append(nic.IPs, (*c.IPAddress)...)
		}
	}

	// 3) 결과 집계
	var list []NICInfo
	for _, nic := range adapters {
		list = append(list, *nic)
	}

	host, _ := os.Hostname()
	return &NetworkInfo{
		Hostname: host,
		Adapters: list,
	}, nil
}
