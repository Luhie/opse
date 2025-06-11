package models

// type AssetPayload struct {
// 	Corporation  string `json:"corporation"`
// 	Department   string `json:"department"`
// 	Team         string `json:"team"`
// 	Name         string `json:"name"`
// 	Position     string `json:"position"`
// 	Location     string `json:"location"`
// 	UserNote     string `json:"userNote"`
// 	Category     string `json:"category"`
// 	ItemType     string `json:"itemType"`
// 	Quantity     int    `json:"quantity"`
// 	Model        string `json:"model"`
// 	Manufacturer string `json:"manufacturer"`
// 	PurchaseDate string `json:"purchaseDate"`
// 	Usage        string `json:"usage"`
// 	AssetNote    string `json:"assetNote"`
// 	AssetName    string `json:"assetName"`
// 	AssetId      string `json:"assetId"`

// 	SystemInfo  SystemInfo   `json:"systemInfo"`
// 	WindowsInfo WindowsInfo  `json:"windowsInfo"`
// 	OfficeList  []OfficeInfo `json:"officeList"`
// }

// type SystemInfo struct {
// 	CPU         CPUInfo         `json:"CPU"`
// 	Motherboard MotherboardInfo `json:"Motherboard"`
// 	RAM         []RAMInfo       `json:"RAM"`
// 	GPU         []GPUInfo       `json:"GPU"`
// 	Disk        []DiskInfo      `json:"Disk"`
// 	Network     []NetworkInfo   `json:"Network"`
// }

// type CPUInfo struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Model        string `json:"model"`
// 	Cores        int    `json:"cores"`
// 	Threads      int    `json:"threads"`
// 	Clock        int    `json:"clock"`
// }

// type MotherboardInfo struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Model        string `json:"model"`
// }

// type RAMInfo struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Model        string `json:"model"`
// 	Capacity     int    `json:"capacity"`
// 	Speed        int    `json:"speed"`
// 	Type         string `json:"type"`
// }

// type GPUInfo struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Model        string `json:"model"`
// 	MemoryMB     int    `json:"memoryMB"`
// }

// type DiskInfo struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Model        string `json:"model"`
// 	CapacityGB   int    `json:"capacityGB"`
// }

// type NetworkInfo struct {
// 	Name string `json:"name"`
// 	MAC  string `json:"mac"`
// 	IP   string `json:"ip"`
// }

type WindowsInfo struct {
	Caption              string `json:"caption"`
	Version              string `json:"version"`
	SerialNumber         string `json:"serial_number"`
	ProductKey           string `json:"product_key"`
	PartialProductKey    string `json:"partial_product_key"`
	Licensed             int    `json:"licensed"`
	ProductChannel       string `json:"product_channel"`
	KmsMachine           string `json:"kms_machine"`
	GracePeriodRemaining int    `json:"grace_period_remaining"`
	EstimatedExpireDate  string `json:"estimated_expire_date"`
	Cracked              int    `json:"cracked"`
}

type OfficeInfo struct {
	Name                string `json:"name"`
	Version             string `json:"version"`
	PartialProductKey   string `json:"partial_product_key"`
	Licensed            int    `json:"licensed"`
	GracePeriod         int    `json:"grace_period"`
	EstimatedExpireDate string `json:"estimated_expire_date"`
	Cracked             int    `json:"cracked"`
}
