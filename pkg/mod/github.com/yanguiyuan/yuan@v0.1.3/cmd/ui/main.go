package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	c := canvas.NewCircle(color.Black)
	//content := container.NewVBox(input, widget.NewButton("Save", func() {
	//	log.Println("Content was:", input.Text)
	//}))

	myWindow.SetContent(c)
	myWindow.ShowAndRun()
}
