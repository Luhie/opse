package software

import (
	"golang.org/x/sys/windows/registry"
)

// Registry 기반 Fallback
func GetOfficeFromRegistry() ([]OfficeInfo, error) {
	keyPath := `SOFTWARE\Microsoft\Office\ClickToRun\Configuration`
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.READ)
	if err != nil {
		return []OfficeInfo{}, nil // 레지스트리에도 없으면 아예 Office 미설치
	}
	defer k.Close()

	productName, _, _ := k.GetStringValue("ProductReleaseIds")
	version, _, _ := k.GetStringValue("ClientVersionToReport")

	info := OfficeInfo{
		Name:                productName,
		Version:             version,
		PartialProductKey:   "(Unknown)",
		Licensed:            -1, // 레지스트리에선 알 수 없음 → 구분 위해 -1 사용
		GracePeriod:         0,
		EstimatedExpireDate: "",
		Cracked:             -1,
	}

	return []OfficeInfo{info}, nil
}
