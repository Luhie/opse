package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

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

	if err := json.Unmarshal([]byte(data), &form); Err != nil {
		return err
	}

	trimAssetForm(&form)

	fmt.Println("테스트 데이터입니다.")
	fmt.Printf(">>> Trim AssetForm: %=v\n", form)
	return nil
}

func (a *App) CreateAsset(data string) error {

	// JSON문자열 언마샬링
	var form AssetForm
	if err := json.Unmarshal([]byte(data), &form); err != nil {
		fmt.Println(">>> CreateAsset Unmarshal Error:", err)
		return err
	}

	// Trim 처리 함수 호출
	trimAssetForm(&form)

	// systemInfo 구조체 다시 JSON []byte로 마샬링
	sysInfoBytes, err := json.Marshal(form.SystemInfo)
	if err != nil {
		fmt.Println(">>> systemInfo Marshal Error:", err)
		return err
	}

	_, err = db.Exec(`
		INSERT INTO assets (
			corporation, department, team, name, position, user_note,
			category, item_type, quantity, model, manufacturer, purchase_date,
			`+"`usage`"+`, asset_note, asset_name, asset_id, system_info
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		form.Corporation, form.Department, form.Team, form.Name, form.Position, form.UserNote,
		form.Category, form.ItemType, form.Quantity, form.Model, form.Manufacturer, form.PurchaseDate,
		form.Usage, form.AssetNote, form.AssetName, form.AssetId, string(sysInfoBytes),
	)

	if err != nil {
		fmt.Print(">>> CreateAsset INSERT Exec Error:", err)
		return err
	}

	fmt.Println(">>> CreateAsset: Success INSERT")
	return err
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
