package views

import (
	"simulador/src/scenes"
    "simulador/src/models"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

type CarObserver struct {
    car *canvas.Image
}

func NewCarObserver(car *models.Car, scene *scenes.ParkingScene) *CarObserver {
    observer := &CarObserver{car: car.GetImage()}
    return observer
}

func (co *CarObserver) Update(pos models.Pos) {
    co.car.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
}
