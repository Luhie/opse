package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"opse/backend/software"
	"opse/backend/system"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type AssetForm struct {
	Corporation  string `json:"corporation"`
	Department   string `json:"department"`
	Team         string `json:"team"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Location     string `json:"location"`
	UserNote     string `json:"userNote"`
	Category     string `json:"category"`
	ItemType     string `json:"itemType"`
	Quantity     int    `json:"quantity"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	PurchaseDate string `json:"purchaseDate"`
	Usage        string `json:"usage"`
	AssetNote    string `json:"assetNote"`
	AssetName    string `json:"assetName"`
	AssetId      string `json:"assetId"`

	SystemInfo struct {
		CPU         system.CPUInfo         `json:"CPU"`
		Motherboard system.MotherboardInfo `json:"Motherboard"`
		RAM         system.RAMInfo         `json:"RAM"`
		GPU         system.GPUs            `json:"GPU"`
		Disk        system.Disks           `json:"Disk"`
		Network     system.NetworkInfo     `json:"Network"`
	} `json:"systemInfo"`

	WindowsInfo software.WindowsInfo  `json:"windowsInfo"` // ✅ 추가
	OfficeList  []software.OfficeInfo `json:"officeList"`  // ✅ 추가
}

/**********************************
* DB전역변수 선언
* InitDB 함수 선언
***********************************************/
var db *sql.DB

func InitDB() error {

	// DSN(Data Source Name) 구성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}

func (a *App) TestAsset(data string) error {
	var form AssetForm

	if err := json.Unmarshal([]byte(data), &form); err != nil {
		return err
	}

	trimAssetForm(&form)

	fmt.Println("테스트 데이터입니다.")
	fmt.Printf(">>> Trim AssetForm: %=v\n", form)
	return nil
}

func (a *App) CreateAsset(data string) error {
	var form AssetForm
	if err := json.Unmarshal([]byte(data), &form); err != nil {
		fmt.Println("▶ CreateAsset Unmarshal 오류:", err)
		return err
	}

	trimAssetForm(&form)

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("▶ 트랜잭션 시작 오류:", err)
		return err
	}

	// 1️⃣ assets 테이블 저장
	res, err := tx.Exec(`
        INSERT INTO assets (
            corporation, department, team, name, position, location, user_note,
            category, item_type, quantity, model, manufacturer, purchase_date,
            `+"`usage`"+`, asset_note, asset_name, asset_id
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		form.Corporation, form.Department, form.Team, form.Name, form.Position,
		form.Location, form.UserNote, form.Category, form.ItemType, form.Quantity,
		form.Model, form.Manufacturer, form.PurchaseDate, form.Usage,
		form.AssetNote, form.AssetName, form.AssetId,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println("▶ assets INSERT 오류:", err)
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2️⃣ system_info JSON 저장 (기존 유지)
	cpuJSON, _ := json.Marshal(form.SystemInfo.CPU)
	mbJSON, _ := json.Marshal(form.SystemInfo.Motherboard)
	ramJSON, _ := json.Marshal(form.SystemInfo.RAM)
	gpuJSON, _ := json.Marshal(form.SystemInfo.GPU)
	diskJSON, _ := json.Marshal(form.SystemInfo.Disk)
	netJSON, _ := json.Marshal(form.SystemInfo.Network)

	_, err = tx.Exec(`
        INSERT INTO system_info (asset_id, cpu, motherboard, ram, gpu, disk, network)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `,
		lastID, string(cpuJSON), string(mbJSON), string(ramJSON),
		string(gpuJSON), string(diskJSON), string(netJSON),
	)
	if err != nil {
		tx.Rollback()
		fmt.Println("▶ system_info INSERT 오류:", err)
		return err
	}

	// 3️⃣ windows_info 저장 추가
	_, err = tx.Exec(`
        INSERT INTO windows_info (
            asset_id, caption, version, serial_number, product_key, partial_product_key, licensed,
            product_channel, kms_machine, grace_period_remaining, estimated_expire_date, cracked
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		lastID,
		form.WindowsInfo.Caption, form.WindowsInfo.Version, form.WindowsInfo.SerialNumber,
		form.WindowsInfo.ProductKey, form.WindowsInfo.PartialProductKey, form.WindowsInfo.Licensed,
		form.WindowsInfo.ProductChannel, form.WindowsInfo.KMSMachine,
		form.WindowsInfo.GracePeriodRemaining, form.WindowsInfo.EstimatedExpireDate, form.WindowsInfo.Cracked,
	)
	if err != nil {
		tx.Rollback()
		fmt.Println("▶ windows_info INSERT 오류:", err)
		return err
	}

	// 4️⃣ office_info 저장 추가
	for _, office := range form.OfficeList {
		_, err = tx.Exec(`
            INSERT INTO office_info (
                asset_id, name, version, partial_product_key, licensed, grace_period, estimated_expire_date, cracked
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        `,
			lastID, office.Name, office.Version, office.PartialProductKey,
			office.Licensed, office.GracePeriod, office.EstimatedExpireDate, office.Cracked,
		)
		if err != nil {
			tx.Rollback()
			fmt.Println("▶ office_info INSERT 오류:", err)
			return err
		}
	}

	// 트랜잭션 커밋
	if err := tx.Commit(); err != nil {
		fmt.Println("▶ 트랜잭션 커밋 오류:", err)
		return err
	}

	fmt.Println("▶ CreateAsset 전체 저장 완료 (asset_id =", lastID, ")")
	return nil
}

// system
func (a *App) GetCPUInfo() (*system.CPUInfo, error) {
	return system.GetCPUInfo()
}
func (a *App) GetMotherboardInfo() (*system.MotherboardInfo, error) {
	return system.GetMotherboardInfo()
}
func (a *App) GetRAMInfo() (*system.RAMInfo, error) {
	return system.GetRAMInfo()
}
func (a *App) GetGPUInfo() (*system.GPUs, error) {
	return system.GetGPUInfo()
}
func (a *App) GetDiskInfo() (*system.Disks, error) {
	return system.GetDiskInfo()
}

// net work
func (a *App) GetNetworkInfo() (*system.NetworkInfo, error) {
	return system.GetNetworkInfo()
}

/* 구조체의 모든 문자열 필드 Trim처리 */
func trimAssetForm(form *AssetForm) {
	form.Corporation = strings.TrimSpace(form.Corporation)
	form.Department = strings.TrimSpace(form.Department)
	form.Team = strings.TrimSpace(form.Team)
	form.Name = strings.TrimSpace(form.Name)
	form.Position = strings.TrimSpace(form.Position)
	form.Location = strings.TrimSpace(form.Location)
	form.UserNote = strings.TrimSpace(form.UserNote)
	form.Category = strings.TrimSpace(form.Category)
	form.ItemType = strings.TrimSpace(form.ItemType)
	form.Model = strings.TrimSpace(form.Model)
	form.Manufacturer = strings.TrimSpace(form.Manufacturer)
	form.PurchaseDate = strings.TrimSpace(form.PurchaseDate)
	form.Usage = strings.TrimSpace(form.Usage)
	form.AssetNote = strings.TrimSpace(form.AssetNote)
	form.AssetName = strings.TrimSpace(form.AssetName)
	form.AssetId = strings.TrimSpace(form.AssetId)

	form.SystemInfo.CPU.Manufacturer = strings.TrimSpace(form.SystemInfo.CPU.Manufacturer)
	form.SystemInfo.CPU.Model = strings.TrimSpace(form.SystemInfo.CPU.Model)

	form.SystemInfo.Motherboard.Manufacturer = strings.TrimSpace(form.SystemInfo.Motherboard.Manufacturer)
	form.SystemInfo.Motherboard.Model = strings.TrimSpace(form.SystemInfo.Motherboard.Model)
	form.SystemInfo.Motherboard.Serial = strings.TrimSpace(form.SystemInfo.Motherboard.Serial)
	form.SystemInfo.Motherboard.Version = strings.TrimSpace(form.SystemInfo.Motherboard.Version)

	// 4) RAMInfo: modules 슬라이스 내부
	for i := range form.SystemInfo.RAM.Modules {
		form.SystemInfo.RAM.Modules[i].Manufacturer = strings.TrimSpace(form.SystemInfo.RAM.Modules[i].Manufacturer)
		form.SystemInfo.RAM.Modules[i].Model = strings.TrimSpace(form.SystemInfo.RAM.Modules[i].Model)
		form.SystemInfo.RAM.Modules[i].DDR = strings.TrimSpace(form.SystemInfo.RAM.Modules[i].DDR)
		// CapacityMB, SpeedMHz 같은 숫자 필드는 생략
	}

	// 5) GPUInfo: cards 슬라이스 내부
	for i := range form.SystemInfo.GPU.Cards {
		form.SystemInfo.GPU.Cards[i].Manufacturer = strings.TrimSpace(form.SystemInfo.GPU.Cards[i].Manufacturer)
		form.SystemInfo.GPU.Cards[i].Model = strings.TrimSpace(form.SystemInfo.GPU.Cards[i].Model)
		form.SystemInfo.GPU.Cards[i].DriverVersion = strings.TrimSpace(form.SystemInfo.GPU.Cards[i].DriverVersion)
		// VRAMMB, Count는 숫자라서 생략
	}

	// 6) DiskInfo: drives 슬라이스 내부
	for i := range form.SystemInfo.Disk.Drives {
		form.SystemInfo.Disk.Drives[i].Manufacturer = strings.TrimSpace(form.SystemInfo.Disk.Drives[i].Manufacturer)
		form.SystemInfo.Disk.Drives[i].Model = strings.TrimSpace(form.SystemInfo.Disk.Drives[i].Model)
		form.SystemInfo.Disk.Drives[i].Serial = strings.TrimSpace(form.SystemInfo.Disk.Drives[i].Serial)
		form.SystemInfo.Disk.Drives[i].Interface = strings.TrimSpace(form.SystemInfo.Disk.Drives[i].Interface)
		form.SystemInfo.Disk.Drives[i].Type = strings.TrimSpace(form.SystemInfo.Disk.Drives[i].Type)
		// SizeGB, Count는 숫자라서 생략
	}

	// 7) NetworkInfo: Hostname 및 adapters 슬라이스 내부
	form.SystemInfo.Network.Hostname = strings.TrimSpace(form.SystemInfo.Network.Hostname)
	for i := range form.SystemInfo.Network.Adapters {
		form.SystemInfo.Network.Adapters[i].Name = strings.TrimSpace(form.SystemInfo.Network.Adapters[i].Name)
		form.SystemInfo.Network.Adapters[i].MAC = strings.TrimSpace(form.SystemInfo.Network.Adapters[i].MAC)
		// IPs가 []string 타입이면, 각 IP 문자열도 Trim 처리
		for j := range form.SystemInfo.Network.Adapters[i].IPs {
			form.SystemInfo.Network.Adapters[i].IPs[j] = strings.TrimSpace(form.SystemInfo.Network.Adapters[i].IPs[j])
		}
		// WiFi(bool)은 건너뜁니다.
	}
}

/*
Software
*/
func (a *App) GetWindowsInfo() (*software.WindowsInfo, error) {
	return software.GetWindowsInfo()
}
func (a *App) GetOfficeInfo() ([]software.OfficeInfo, error) {
	return software.GetOfficeInfo()
}
func (a *App) GetHwpInfo() (*software.HwpInfo, error) {
	return software.GetHwpInfo()
}

// Get List
func (a *App) GetAssetList() ([]system.Asset, error) {
	return system.GetAssetList()
}
