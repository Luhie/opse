package system

type Asset struct {
	AssetID      string `json:"assetId"`
	AssetName    string `json:"assetName"`
	ItemType     string `json:"itemType"`
	Manufacturer string `json:"manufacturer"`
	PurchaseDate string `json:"purchaseDate"`
}

func GetAssetList() ([]Asset, error) {
	// 임시 더미 데이터 - 추후 DB 연동으로 교체 예정
	assets := []Asset{
		{AssetID: "A001", AssetName: "노트북", ItemType: "PC본체", Manufacturer: "LG", PurchaseDate: "2024-12-01"},
		{AssetID: "A002", AssetName: "모니터", ItemType: "모니터", Manufacturer: "Samsung", PurchaseDate: "2025-01-10"},
	}
	return assets, nil
}
