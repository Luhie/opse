package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

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
}

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

func (a *App) CreateAsset(data string) error {
	var form AssetForm
	if err := json.Unmarshal([]byte(data), &form); err != nil {
		return err
	}

	_, err := db.Exec(`
		INSERT INTO assets (
			corporation, department, team, name, position, user_note,
			category, item_type, quantity, model, manufacturer, purchase_date,
			`+"`usage`"+`, asset_note, asset_name, asset_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		form.Corporation, form.Department, form.Team, form.Name, form.Position, form.UserNote,
		form.Category, form.ItemType, form.Quantity, form.Model, form.Manufacturer, form.PurchaseDate,
		form.Usage, form.AssetNote, form.AssetName, form.AssetId,
	)

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
