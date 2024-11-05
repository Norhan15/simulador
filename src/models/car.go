package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Car struct {
	carImage *canvas.Image
	CarPosX  float32
	CarPosY  float32 
}

func NewCar(imagePath string, initialX, initialY float32) *Car {
	car := &Car{
		carImage: canvas.NewImageFromURI(storage.NewFileURI(imagePath)),
		CarPosX:  initialX, 
		CarPosY:  initialY,
	}
	car.carImage.Resize(fyne.NewSize(50, 30))
	return car
}

func (c *Car) GetImage() *canvas.Image {
	return c.carImage
}
