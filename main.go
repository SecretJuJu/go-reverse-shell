package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-reverse-shell/speedtest"
)

type MyApp struct {
	app                 *fyne.App
	mainWindow          *fyne.Window
	label               *widget.Label
	downloadResultLabel *widget.Label
	uploadResultLabel   *widget.Label
}

func initApp() *MyApp {
	// 앱 생성
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Speed Test")

	// 레이블 추가
	label := widget.NewLabel("Speed Test")

	// resultLabel 추가
	downloadResultLabel := widget.NewLabel("Download speed: 0.00 Mbps")
	downloadResultLabel.TextStyle.Bold = true
	uploadResultLabel := widget.NewLabel("Upload speed: 0.00 Mbps")
	uploadResultLabel.TextStyle.Bold = true

	return &MyApp{
		app:                 &myApp,
		mainWindow:          &myWindow,
		label:               label,
		downloadResultLabel: downloadResultLabel,
		uploadResultLabel:   uploadResultLabel,
	}
}

func main() {
	app := initApp()
	testChannel := make(chan speedtest.TestResult)

	myApp := *app.app
	window := *app.mainWindow
	label := app.label
	downloadResultLabel := app.downloadResultLabel
	uploadResultLabel := app.uploadResultLabel

	// 버튼, 레이블 생성
	window.SetContent(container.NewVBox(
		label,
		downloadResultLabel,
		uploadResultLabel,
		widget.NewButton("Test", func() {
			label.SetText("Testing...")
			// Test 함수 호출, 결과를 채널을 통해 전달
			go speedtest.Test(testChannel)
			go handleTestResult(testChannel, app)
		}),
		widget.NewButton("Quit", func() {
			myApp.Quit()
		}),
	))

	// 창 크기 설정
	window.Resize(fyne.NewSize(400, 300))

	// 창 표시
	window.ShowAndRun()
}

func handleTestResult(resultChannel chan speedtest.TestResult, app *MyApp) {
	result := <-resultChannel
	app.label.SetText("done")

	downloadSpeedMbps := result.DownloadSpeed / 1_000_000
	uploadSpeedMbps := result.UploadSpeed / 1_000_000

	app.downloadResultLabel.SetText(fmt.Sprintf("Download speed: %.2f Mbps", downloadSpeedMbps))
	app.uploadResultLabel.SetText(fmt.Sprintf("Upload speed: %.2f Mbps", uploadSpeedMbps))
}
