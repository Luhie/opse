package software

import (
	"golang.org/x/sys/windows/registry"
)

// 한컴(HWP+Office) 통합 정보 구조체
type HwpInfo struct {
	ProductName string `json:"product_name"`
	Version     string `json:"version"`
	InstallPath string `json:"install_path"`
}

// 메인 호출 함수
func GetHwpInfo() (*HwpInfo, error) {
	// 1. 구버전 HWP 경로 시도
	if info, err := readHwpRegistry(`SOFTWARE\HNC\Hwp`); err == nil && info != nil {
		return info, nil
	}

	if info, err := readHwpRegistry(`SOFTWARE\WOW6432Node\HNC\Hwp`); err == nil && info != nil {
		return info, nil
	}

	// 2. 신버전 Office 경로 시도
	if info, err := readOfficeRegistry(`SOFTWARE\HNC\Office`); err == nil && info != nil {
		return info, nil
	}

	if info, err := readOfficeRegistry(`SOFTWARE\WOW6432Node\HNC\Office`); err == nil && info != nil {
		return info, nil
	}

	// 모두 실패시 설치 안됨으로 간주
	return nil, nil
}

// 구형 HWP 경로 레지스트리 읽기
func readHwpRegistry(path string) (*HwpInfo, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.READ)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	productName, _, _ := k.GetStringValue("ProductName")
	version, _, _ := k.GetStringValue("Version")
	installPath, _, _ := k.GetStringValue("InstallPath")

	return &HwpInfo{
		ProductName: productName,
		Version:     version,
		InstallPath: installPath,
	}, nil
}

// 신형 Office 경로 레지스트리 읽기
func readOfficeRegistry(path string) (*HwpInfo, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.READ)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	version, _, _ := k.GetStringValue("Version")
	installPath, _, _ := k.GetStringValue("InstallPath")

	return &HwpInfo{
		ProductName: "한컴오피스",
		Version:     version,
		InstallPath: installPath,
	}, nil
}
