package main

import (
	"embed"
	"log"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/lxn/win"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 환경변수 로드
	if err := godotenv.Load(); err != nil {
		log.Fatal("환경변수 로딩 실패: .env 파일이 없습니다.")
	}

	// Create an instance of the app structure
	app := NewApp()

	if err := InitDB(); err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	// width, height := getScreenSize()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "opse",
		// Width:  width,
		// Height: height,
		Width:  1024,
		Height: 768,
		// MinWidth:  1280,
		// MinHeight: 720,
		// Fullscreen: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func getScreenSize() (width int, height int) {
	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	return screenWidth, screenHeight
}
