package scenes

import (
	"image/color"
	"simulador/src/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type ParkingScene struct {
	window       fyne.Window
	container    *fyne.Container
	Car          *models.Car // Campo exportable
	parkingSpots []*fyne.Container
	isMoving     bool
	step         float32
	endPosX      float32
	endPosY      float32
	observers    []models.Observer
}

func NewParkingScene(window fyne.Window) *ParkingScene {
	return &ParkingScene{window: window}
}

func (ps *ParkingScene) Init() {
	ps.container = container.NewWithoutLayout()

	background := canvas.NewRectangle(color.RGBA{R: 180, G: 180, B: 180, A: 255})
	background.Resize(fyne.NewSize(870, 580))
	ps.container.Add(background)

	ps.Car = models.NewCar("./assets/car.png", 10, 580-30) 
	ps.container.Add(ps.Car.GetImage())

	ps.Car.GetImage().Move(fyne.NewPos(ps.Car.CarPosX, ps.Car.CarPosY))

	ps.endPosX = 750
	ps.endPosY = 30
	ps.step = 5
	ps.isMoving = true

	ps.startAnimation()

	parkingWidth := 100
	parkingHeight := 50
	offsetX := 50
	offsetY := 50
	spaceBetweenSlots := 50

	for row := 0; row < 4; row++ {
		for col := 0; col < 5; col++ {
			slot := canvas.NewRectangle(color.RGBA{R: 0, G: 100, B: 0, A: 255})
			slot.Resize(fyne.NewSize(float32(parkingWidth), float32(parkingHeight)))
			slotPosX := float32(offsetX + col*(parkingWidth+spaceBetweenSlots))
			slotPosY := float32(offsetY + row*(parkingHeight+spaceBetweenSlots))
			slot.Move(fyne.NewPos(slotPosX, slotPosY))
			ps.container.Add(slot)
			ps.parkingSpots = append(ps.parkingSpots, container.NewWithoutLayout(slot))
		}
	}

	entrance := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	entrance.Resize(fyne.NewSize(750, 30))
	entrance.Move(fyne.NewPos(10, float32(offsetY+4*(parkingHeight+spaceBetweenSlots))))
	ps.container.Add(entrance)

	ps.window.SetContent(ps.container)
}

func (ps *ParkingScene) startAnimation() {
	go func() {
		ps.isMoving = true
		movingRight := true // Fase 1: Movimiento hacia la derecha

		for ps.isMoving {
			time.Sleep(16 * time.Millisecond)

			// Fase 1: Movimiento horizontal hacia la derecha
			if movingRight {
				if ps.Car.CarPosX < ps.endPosX {
					ps.Car.CarPosX += ps.step
				} else {
					// Notificar a los observadores en la primera detención
					ps.NotifyAll(models.Pos{
						X: int32(ps.Car.CarPosX), // Convertir a int32
						Y: int32(ps.Car.CarPosY), // Convertir a int32
					})

					// Cambiar a la fase 2: movimiento hacia arriba
					movingRight = false
				}
			}

			// Fase 2: Movimiento vertical hacia arriba
			if !movingRight {
				if ps.Car.CarPosY > ps.endPosY {
					ps.Car.CarPosY -= ps.step
				} else {
					ps.isMoving = false
				}
			}

			// Actualizar posición gráfica
			ps.Car.GetImage().Move(fyne.NewPos(ps.Car.CarPosX, ps.Car.CarPosY))
			ps.container.Refresh()
		}
	}()
}

// Registrar observadores
func (ps *ParkingScene) Register(observer models.Observer) {
	ps.observers = append(ps.observers, observer)
}

func (ps *ParkingScene) Unregister(observer models.Observer) {
	for i, obs := range ps.observers {
		if obs == observer {
			ps.observers = append(ps.observers[:i], ps.observers[i+1:]...) // Eliminar el observador
			break
		}
	}
}

// Notificar a todos los observadores
func (ps *ParkingScene) NotifyAll(pos models.Pos) {
	for _, observer := range ps.observers {
		observer.Update(pos)
	}
}
