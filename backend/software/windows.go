package software

import (
	"errors"
	"time"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

type WindowsInfo struct {
	Caption              string `json:"caption"`
	Version              string `json:"version"`
	SerialNumber         string `json:"serial_number"`
	ProductKey           string `json:"product_key"`
	PartialProductKey    string `json:"partial_product_key"`
	Licensed             int    `json:"licensed"` // bool -> int (0/1)
	ProductChannel       string `json:"product_channel"`
	KMSMachine           string `json:"kms_machine"`
	GracePeriodRemaining uint32 `json:"grace_period_remaining"`
	EstimatedExpireDate  string `json:"estimated_expire_date"`
	Cracked              int    `json:"cracked"` // 크랙 의심 여부 (0/1)
}

func GetWindowsBasicInfo() (*WindowsInfo, error) {
	var dst []struct {
		Caption      string
		Version      string
		SerialNumber string
	}
	err := wmi.Query("SELECT Caption, Version, SerialNumber FROM Win32_OperatingSystem", &dst)
	if err != nil || len(dst) == 0 {
		return nil, err
	}

	info := &WindowsInfo{
		Caption:      dst[0].Caption,
		Version:      dst[0].Version,
		SerialNumber: dst[0].SerialNumber,
	}

	return info, nil
}

func GetWindowsLicenseDetail(info *WindowsInfo) {
	var products []struct {
		LicenseStatus        uint32
		Name                 string
		PartialProductKey    string
		ProductKeyChannel    string
		GracePeriodRemaining uint32
	}
	err := wmi.Query("SELECT LicenseStatus, Name, PartialProductKey, ProductKeyChannel, GracePeriodRemaining FROM SoftwareLicensingProduct", &products)
	if err != nil {
		info.Licensed = 0
		info.PartialProductKey = "Unknown"
		info.ProductChannel = "Unknown"
		info.GracePeriodRemaining = 0
		info.EstimatedExpireDate = "Unknown"
		info.Cracked = 0
		return
	}

	for _, p := range products {
		if p.LicenseStatus == 1 {
			info.Licensed = 1
			info.PartialProductKey = p.PartialProductKey
			info.ProductChannel = p.ProductKeyChannel
			info.GracePeriodRemaining = p.GracePeriodRemaining
			info.EstimatedExpireDate = calculateExpireDate(p.GracePeriodRemaining)
			info.Cracked = DetectCracked(info)
			break
		}
	}
}

func calculateExpireDate(gracePeriod uint32) string {
	now := time.Now()
	expire := now.Add(time.Duration(gracePeriod) * 24 * time.Hour)
	return expire.Format("2006-01-02")
}

func GetKMSMachine(info *WindowsInfo) {
	var service []struct {
		KeyManagementServiceMachine string
	}
	err := wmi.Query("SELECT KeyManagementServiceMachine FROM SoftwareLicensingService", &service)
	if err != nil {
		info.KMSMachine = "Unknown"
		return
	}

	if len(service) > 0 {
		info.KMSMachine = service[0].KeyManagementServiceMachine
	} else {
		info.KMSMachine = ""
	}
}

func ReadWindowsProductKey() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.READ)
	if err != nil {
		return "", err
	}
	defer key.Close()

	digitalProductId, _, err := key.GetBinaryValue("DigitalProductId")
	if err != nil {
		return "", err
	}

	productKey, err := decodeProductKey(digitalProductId)
	if err != nil {
		return "", err
	}
	return productKey, nil
}

func decodeProductKey(digitalProductId []byte) (string, error) {
	if len(digitalProductId) < 67 {
		return "", errors.New("invalid digitalProductId length")
	}

	const keyOffset = 52
	chars := "BCDFGHJKMPQRTVWXY2346789"
	decodedKey := make([]byte, 29)

	for i := 28; i >= 0; i-- {
		if (i+1)%6 == 0 {
			decodedKey[i] = '-'
		} else {
			accumulator := 0
			for j := 14; j >= 0; j-- {
				accumulator = accumulator*256 + int(digitalProductId[keyOffset+j])
				digitalProductId[keyOffset+j] = byte(accumulator / 24)
				accumulator = accumulator % 24
			}
			decodedKey[i] = chars[accumulator]
		}
	}

	return string(decodedKey), nil
}

func GetWindowsInfo() (*WindowsInfo, error) {
	info, err := GetWindowsBasicInfo()
	if err != nil {
		return nil, err
	}

	productKey, err := ReadWindowsProductKey()
	if err == nil {
		info.ProductKey = productKey
	} else {
		info.ProductKey = "Unknown"
	}

	GetWindowsLicenseDetail(info)
	GetKMSMachine(info)

	return info, nil
}

// 🧠 핵심 자동 진단 로직:
func DetectCracked(info *WindowsInfo) int {
	if info.ProductChannel == "Volume:GVLK" &&
		(info.GracePeriodRemaining > 1000 || info.KMSMachine == "0.0.0.0") {
		return 1
	}
	return 0
}
