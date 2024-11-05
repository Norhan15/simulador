package main

import (
	"simulador/src/scenes"
	"simulador/src/views"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Parking Scene")

	parkingScene := scenes.NewParkingScene(myWindow)
	parkingScene.Init()

	carObserver := views.NewCarObserver(parkingScene.Car, parkingScene)

    // Registrar el observador
    parkingScene.Register(carObserver)

	myWindow.Resize(fyne.NewSize(900, 600))
	myWindow.ShowAndRun()
}
