package software

import (
	"github.com/StackExchange/wmi"
)

type OfficeWMIProduct struct {
	Name                 string
	Version              string
	PartialProductKey    string
	LicenseStatus        uint32
	GracePeriodRemaining uint32
}

type OfficeInfo struct {
	Name                string `json:"name"`
	Version             string `json:"version"`
	PartialProductKey   string `json:"partial_product_key"`
	Licensed            int    `json:"licensed"`
	GracePeriod         uint32 `json:"grace_period"`
	EstimatedExpireDate string `json:"estimated_expire_date"`
	Cracked             int    `json:"cracked"`
}

func GetOfficeInfo() ([]OfficeInfo, error) {
	// 먼저 WMI 탐지 시도
	products, err := getOfficeFromWMI()
	if err != nil || len(products) == 0 {
		// WMI 실패시 → 레지스트리 fallback 시도
		return GetOfficeFromRegistry()
	}
	return products, nil
}

func getOfficeFromWMI() ([]OfficeInfo, error) {
	var products []OfficeWMIProduct

	err := wmi.Query("SELECT Name, Version, LicenseStatus, PartialProductKey, GracePeriodRemaining FROM OfficeSoftwareProtectionProduct", &products)
	if err != nil {
		return nil, err
	}

	var results []OfficeInfo

	for _, p := range products {
		info := OfficeInfo{
			Name:                p.Name,
			Version:             p.Version,
			PartialProductKey:   p.PartialProductKey,
			Licensed:            convertLicense(p.LicenseStatus),
			GracePeriod:         p.GracePeriodRemaining,
			EstimatedExpireDate: CalculateExpireDate(p.GracePeriodRemaining),
			Cracked:             detectCracked(p),
		}
		results = append(results, info)
	}

	return results, nil
}

func convertLicense(status uint32) int {
	if status == 1 {
		return 1
	}
	return 0
}

func detectCracked(p OfficeWMIProduct) int {
	if p.LicenseStatus == 1 && p.GracePeriodRemaining > 1000 {
		return 1
	}
	return 0
}
