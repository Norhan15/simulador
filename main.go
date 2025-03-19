package main

import (
	"fmt"

    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Ejemplo Completo")
    myWindow.Resize(fyne.NewSize(400, 300))

    // Etiqueta para mensajes
    messageLabel := widget.NewLabel("Seleccione una opción:")

    // Radio Group
    radio := widget.NewRadioGroup([]string{"Opción A", "Opción B"}, func(selected string) {
        messageLabel.SetText(fmt.Sprintf("Seleccionaste: %s", selected))
    })

    // Slider
    sliderValue := widget.NewLabel("Valor del slider: 0.0")
    slider := widget.NewSlider(0, 10)
    slider.OnChanged = func(value float64) {
        sliderValue.SetText(fmt.Sprintf("Valor del slider: %.1f", value))
    }

    // Botón con mensaje
    resultLabel := widget.NewLabel("")
    button := widget.NewButton("Enviar", func() {
        if radio.Selected == "" {
            resultLabel.SetText("Por favor, selecciona una opción.")
        } else {
            resultLabel.SetText(fmt.Sprintf("Enviado: %s con valor %.1f", radio.Selected, slider.Value))
        }
    })

    // Contenedor
    content := container.NewVBox(
        messageLabel,
        radio,
        slider,
        sliderValue,
        button,
        resultLabel,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
